package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"eglass.com/utils"

	"github.com/coreos/etcd/clientv3"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	GET    = "GET"
	POST   = "POST"
	Prefix = "/api/v3/"
)

type ServerInfo struct {
	ID          string `json:"id"`   // 服务器ID
	IP          string `json:"ip"`   // 对外连接服务的 IP
	Port        string `json:"port"` // 对外服务端口，本机或者端口映射后得到的
	Host        string `json:"host"`
	Address     string `json:"address"`
	ServiceName string `json:"serviceName"`
}

var (
	// Etcd = os.Getenv("Etcd") //"t.epeijing.cn:2379"
	Etcd         = "t.epeijing.cn:2379"
	cli, etcdErr = clientv3.New(clientv3.Config{
		Endpoints:   []string{Etcd},
		DialTimeout: 5 * time.Second,
	})
	remoteRequest  = utils.InitRequest()
	noServiceError = errors.New("no service")
	services       = NewServices()
)

type Services struct {
	lock        sync.RWMutex
	idNameMap   map[string]string
	servicesMap map[string][]ServerInfo
}

func (s *Services) Init(cli *clientv3.Client) error {
	resp, err := cli.Get(context.Background(), "service-", clientv3.WithPrefix())
	if err != nil {
		return err
	}
	for _, kv := range resp.Kvs {
		var serverInfo ServerInfo
		err := json.Unmarshal(kv.Value, &serverInfo)
		if err != nil {
			log.Println("etcd registered format error", err)
			continue
		}
		s.addServiceInstance(serverInfo)
	}
	log.Println("services init finished")
	return nil
}

// 为服务列表添加一个服务实例
func (s *Services) addServiceInstance(info ServerInfo) {
	serviceName := info.ServiceName
	s.lock.Lock()
	log.Println("add serviceInstance to servicesMap, info:", info)
	if s.servicesMap[serviceName] == nil {
		s.servicesMap[serviceName] = []ServerInfo{info}
	} else {
		s.servicesMap[serviceName] = append(s.servicesMap[serviceName], info)
	}
	s.idNameMap[info.ID] = serviceName
	s.lock.Unlock()
}

// 在服务列表中删除一个服务实例
func (s *Services) removeServiceInstance(serviceName, ID string) {
	instances := s.servicesMap[serviceName]
	for index, instance := range instances {
		if instance.ID == ID {
			s.lock.Lock()
			log.Println("remove serviceInstance from servicesMap")
			delete(s.idNameMap, ID)
			s.servicesMap[serviceName] = append(instances[:index], instances[index+1:]...)
			log.Println("serviceName:", serviceName, "  the latest serviceInstances:", s.servicesMap[serviceName])
			s.lock.Unlock()
			break
		}
	}
}

func NewServices() *Services {
	return &Services{
		lock:        sync.RWMutex{},
		idNameMap:   make(map[string]string),
		servicesMap: make(map[string][]ServerInfo),
	}
}

func main() {

	if etcdErr != nil {
		log.Panicf("connect to etcd error: %v", etcdErr)
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Any(Prefix+":service/*", handleAll)
	e.Logger.Fatal(e.Start(":3030"))
}

func handleAll(c echo.Context) error {
	//解析服务
	url := c.Request().RequestURI
	serviceName := c.Param("service")
	serviceEndpoint, err := getService(serviceName)
	if err != nil {
		c.String(503, "")
	}
	skip := len(Prefix) + len(serviceName)
	serviceFullURL := serviceEndpoint + url[skip:]
	fmt.Println("serviceURL:", serviceFullURL)
	if c.QueryString() != "" {
		serviceFullURL += c.QueryString()
	}
	req := c.Request()
	var proxyReq *http.Request
	var reqError error
	if req.Method == "GET" {
		proxyReq, reqError = http.NewRequest(req.Method, serviceFullURL, nil)
		if reqError != nil {
			log.Printf("can not make [GET] request, %v", reqError)
			return reqError
		}
	} else {
		defer req.Body.Close()
		body, error := ioutil.ReadAll(req.Body)
		if reqError != nil {
			return error
		}
		reqBody := bytes.NewReader(body)
		proxyReq, reqError = http.NewRequest(req.Method, serviceFullURL, reqBody)
		if reqError != nil {
			log.Printf("can not make request with body, %v", err)
			return reqError
		}
	}
	proxyReq.Header = req.Header
	log.Printf("proxying request to url: %s", serviceFullURL)
	resp, error := remoteRequest.GetClient().Do(proxyReq)
	if error != nil {
		return error
	}
	defer resp.Body.Close()
	_, err = io.Copy(c.Response(), resp.Body)
	return error
}

func getService(serviceName string) (string, error) {
	servicesInfos := services.servicesMap[serviceName]
	services.lock.RLock()
	total := len(servicesInfos)
	services.lock.RUnlock()
	if total == 0 {
		return "", noServiceError
	}
	if total == 1 {
		return "http://" + servicesInfos[0].Address, nil
	}
	log.Printf(" %d available ", total)
	rand.Seed(time.Now().UTC().UnixNano())
	randIndex := rand.Intn(total)
	log.Printf("use index: %d", randIndex)
	return "http://" + servicesInfos[randIndex].Address, nil
}
