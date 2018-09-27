package microservices

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Service struct {
	Server *echo.Echo
	Pes    *Pes
}
type Pes struct {
	Redis string `json:"redis"`
	Mysql string `json:"mysql"`
}

type EglassHandlerFunc func(e echo.Context, pes *Pes) error

func (scrm *Service) ToEcho(f EglassHandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		return f(e, scrm.Pes)
	}
}
func (scrm *Service) Add(method, path string, handler EglassHandlerFunc) {
	scrm.Server.Add(method, path, scrm.ToEcho(handler))
}
func (s *Service) Start(port string) {
	s.Server.Logger.Fatal(s.Server.Start(port))
}

func Init() *Service {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return &Service{
		Server: e,
		Pes:    &Pes{"redis", "mysql"},
	}
}

type MicroService interface {
	Port() string
	Start(s *grpc.Server)
	Name() string
}

func RegisterService(m MicroService) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return err
	}
	resp, err := cli.Grant(context.Background(), 10)
	defer cli.Close()
	if err != nil {
		return err
	}
	_, err = cli.Lease.KeepAliveOnce(context.Background(), resp.ID)
	if err != nil {
		return err
	}
	serviceKey, serviceValue := "/services/"+m.Name(), "localhost"+m.Port()
	_, err = cli.Put(context.Background(), serviceKey, serviceValue, clientv3.WithLease(resp.ID), clientv3.WithCreatedNotify())
	if err != nil {
		return err
	}
	ticker := time.Tick(8 * time.Second)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-ticker:
			_, err := cli.Lease.KeepAliveOnce(context.Background(), resp.ID)
			if err != nil {
				log.Println("keepalived failed")
				return err
			}
			log.Printf("keepalived")
		case <-c:
			_, err := cli.KV.Delete(context.Background(), serviceKey)
			if err != nil {
				log.Printf("fail to delete service: %s\n", m.Name())
			}
			log.Fatal("exit")
		}
	}
}

func StartGrpcServer(m MicroService) {
	lis, err := net.Listen("tcp", m.Port())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	m.Start(s)
	reflection.Register(s)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		log.Printf("micro service started at %s", m.Port())
		if err := s.Serve(lis); err != nil {
			wg.Done()
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	go func() {
		defer wg.Done()
		error := RegisterService(m)
		if error != nil {
			log.Printf("service registry error: %v ", error)
		}
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-c:
			log.Fatal("exit")
		}
	}()

	wg.Wait()
}
