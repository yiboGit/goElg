package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"eglass.com/brisk"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

type UserServerClient struct {
	lock             sync.RWMutex
	serviceInstances []brisk.ServerInfo
}

type RequestIDQueue struct {
	requestIDs []string
	lock       sync.RWMutex
}

var (
	uClient = UserServerClient{
		lock:             sync.RWMutex{},
		serviceInstances: []brisk.ServerInfo{},
	}
	requestIDQueue = RequestIDQueue{
		lock:       sync.RWMutex{},
		requestIDs: []string{},
	}
)

func (uClient *UserServerClient) Init() error {
	log.Printf("getaddress")
	serviceInstances, err := brisk.GetServiceAddress("u_service")
	log.Printf("get over")
	if err != nil {
		return err
	}
	log.Printf("1")
	uClient.serviceInstances = serviceInstances
	go func() {
		// 监听etcd中服务实例变化
		log.Printf("2")
		cli, etcdErr := clientv3.New(clientv3.Config{
			Endpoints:   []string{"t.epeijing.cn:2379"},
			DialTimeout: 5 * time.Second,
		})
		if etcdErr != nil {
			log.Panicf("connect to etcd error: %v", etcdErr)
		}
		watcher := clientv3.NewWatcher(cli)
		w := watcher.Watch(context.Background(), "service-u_service-", clientv3.WithPrefix())
		log.Printf("3")
		for {
			select {
			case watchResp := <-w:
				for _, event := range watchResp.Events {
					if event.Type == mvccpb.PUT {
						var server brisk.ServerInfo
						err := json.Unmarshal(event.Kv.Value, &server)
						if err != nil {
							log.Println("etcd registered format error", err)
						}
						uClient.lock.Lock()
						exist, _ := uClient.isExist(server.ID)
						if !exist {
							uClient.serviceInstances = append(uClient.serviceInstances, server)
						}
						uClient.lock.Unlock()
					} else if event.Type == mvccpb.DELETE {
						key := string(event.Kv.Key)
						nameID := strings.Split(key[len("service-"):], "-")
						uClient.lock.Lock()
						exist, index := uClient.isExist(nameID[1])
						if exist {
							uClient.serviceInstances = append(uClient.serviceInstances[:index], uClient.serviceInstances[index+1:]...)
						}
						uClient.lock.Unlock()
					}
				}
			}
		}
	}()

	return nil
}

func (uClient *UserServerClient) isExist(ID string) (bool, int) {
	exist := false
	var i int
	for index, item := range uClient.serviceInstances {
		if item.ID == ID {
			exist = true
			i = index
			break
		}
	}
	return exist, i
}

func (uClient *UserServerClient) GetUser(p UserReq) UserResp {
	var result UserResp
	err := uClient.request(p, "GetUser", &result)
	if err != nil {
		result.Err.Message = err.Error()
		return result
	}
	return result
}
func (uClient *UserServerClient) AddUser(p UserReq) UserResp {
	var result UserResp
	err := uClient.request(p, "AddUser", &result)
	if err != nil {
		log.Println("有错")
		result.Err.Message = err.Error()
		return result
	}
	log.Println("有错")
	return result
}

func (uClient *UserServerClient) request(param interface{}, method string, respResult interface{}) error {
	// requestId := context.QueryParam("requestId")
	// if requestId == "" {
	// 	requestId = fmt.Sprintf("%v", xid.New())
	// }
	// // requestId = fmt.Sprintf("?requestId=%s", requestId)
	// requestIDQueue.lock.Lock()
	// requestId := requestIDQueue.requestIDs[0]
	// requestIDQueue.lock.Unlock()


	n := len(uClient.serviceInstances)
	if n == 0 {
		return errors.New("service instance is empty")
	}
	var b bytes.Buffer
	u, _ := json.Marshal(param)
	b.Write(u)
	//TODO  这里做选择 例如负载均衡 暂时随机数
	r := rand.New(rand.NewSource(time.Now().Unix()))
	i := r.Intn(n)
	address := uClient.serviceInstances[i].Address
	url := "http://" + address + "/" + "UserServer" + "/" + method
	resp, err := http.Post(url, "application/json", &b)
	if err != nil {
		return err
	}
	b.ReadFrom(resp.Body)
	json.Unmarshal(b.Bytes(), respResult)
	return nil
}
