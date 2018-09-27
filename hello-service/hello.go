package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	// "github.com/satori/go.uuid"
	"github.com/rs/xid"

	"github.com/coreos/etcd/clientv3"
)

type Serverinfo struct {
	ID          string `json:"id"`   // 服务器ID
	IP          string `json:"ip"`   // 对外连接服务的 IP
	Port        string `json:"port"` // 对外服务端口，本机或者端口映射后得到的
	Host        string `json:"host"`
	Address     string `json:"address"`
	ServiceName string `json:"serviceName"`
}

var (
	IP           = os.Getenv("IP")
	Port         = os.Getenv("Port")
	Host         = os.Getenv("Host")
	Etcd         = os.Getenv("Etcd") //"t.epeijing.cn:2379"
	ServiceName  = os.Getenv("ServiceName")
	cli, etcdErr = clientv3.New(clientv3.Config{
		Endpoints:   []string{Etcd},
		DialTimeout: 5 * time.Second,
	})
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s:%s-%s", Host, Port, "Hello update")
			fmt.Fprintf(w, "Hello update")
		})
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second)

		if etcdErr != nil {
			log.Fatal("Error: cannot connect to etcd ", etcdErr)
		}
		guid := xid.New()
		xID := fmt.Sprintf("%s", guid)
		key := fmt.Sprintf("%s-%s-%s", "service", ServiceName, xID)
		address := Host + ":" + Port
		serviceInfo := Serverinfo{
			ID:          xID,
			IP:          IP,
			Port:        Port,
			Host:        Host,
			Address:     address,
			ServiceName: ServiceName,
		}
		log.Printf("key : %s ; value : %v", key, serviceInfo)
		value, err := json.Marshal(serviceInfo)
		if err != nil {
			log.Fatal("Error: Service Info has Error", err)
		}
		//首次注册服务
		serverRegiste(key, value)
		//设置心跳时间，准备注册服务
		ticker := time.NewTicker(10 * time.Second)
		c := make(chan os.Signal, 1)
		//监测os的三种关于退出，销毁的信号
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		for {
			select {
			case <-ticker.C:
				//定时注册,创建10s的租期，意为每10s发送一次心跳
				serverRegiste(key, value)
			case <-c:
				//退出时（销毁时），自动解除注册
				cli.Delete(context.Background(), key)
				log.Println("service unregiste finish")
				return
			}
		}
	}()
	wg.Wait()
}

func serverRegiste(key string, value []byte) {
	lease, err := cli.Grant(context.Background(), 10)
	if err != nil {
		log.Fatal("Error: etcd create lease has error", err)
	}
	_, err = cli.Put(context.Background(), key, string(value), clientv3.WithLease(lease.ID))
	if err != nil {
		log.Fatal("Error: service info registe error", err)
	}
}
