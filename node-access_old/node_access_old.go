package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"eglass.com/brisk"
	"github.com/coreos/etcd/clientv3"
	"github.com/labstack/echo"
)

const (
	urlPrefix = "/api/register/"
	// 操作类型是put
	PUT = "put"
	// 操作类型是delete
	DEL = "delete"
)

var (
	// Etcd         = os.Getenv("Etcd")
	Etcd         = "t.epeijing.cn:2379"
	cli, etcdErr = clientv3.New(clientv3.Config{
		Endpoints:   []string{Etcd},
		DialTimeout: 5 * time.Second,
	})
)

func main() {
	if etcdErr != nil {
		log.Fatal("Error(node_access): cannot connect to etcd ", etcdErr)
	}
	log.Println("Successful(node_access): connect to etcd")

	e := echo.New()
	e.POST(urlPrefix+":action/:key", handleAll)
	e.Logger.Fatal(e.Start(":9999"))
}

func handleAll(c echo.Context) error {
	// 解析服务
	action := c.Param("action")
	key := c.Param("key")
	log.Printf("Info(node_access): action: %s, key: %s \n", action, key)
	if action == PUT {
		log.Printf("Info(node_access): handle %s \n", action)
		var serverInfo brisk.ServerInfo
		c.Bind(&serverInfo)
		log.Printf("Info(node_access): value %v \n", serverInfo)
		value, err := json.Marshal(serverInfo)
		if err != nil {
			log.Printf("Error(node_access): Service Info has Error, %v \n", err)
			return err
		}
		err = serverRegister(cli, key, value)
		if err != nil {
			log.Printf("%v \n", err)
			return err
		}
	} else if action == DEL {
		log.Println("delete start")
		var rmMsg map[string]string
		c.Bind(&rmMsg)
		log.Printf("delete message : %v \n", rmMsg)
		value, err := json.Marshal(rmMsg)
		if err != nil {
			log.Printf("RM-Error(node_access):RM value Marshal error %v \n", err)
			return err
		}
		log.Printf("delete put value \n")
		_, err = cli.Put(context.Background(), key, string(value))
		if err != nil {
			log.Printf("RM-Error(node_access): put remove info error, %v", err)
			return err
		}
		log.Printf("RM-info(node_access): put remove info successfully \n")
		array := strings.Split(key, "-")
		if len(array) == 0 {
			err := errors.New("key length error")
			log.Printf("RM-Error(node_access): RM-key message error, %v", err)
			return err
		}

		serviceKey := fmt.Sprintf("service-%s-%s", rmMsg["serviceName"], array[2])
		cli.Delete(context.Background(), serviceKey)
		log.Println("service unregister finish (node_access)")
	}
	c.String(200, "accepted")
	return nil
}

func serverRegister(cli *clientv3.Client, key string, value []byte) error {
	log.Println("register start")
	lease, err := cli.Grant(context.Background(), 10)
	if err != nil {
		log.Println("lease error")
		msg := fmt.Sprintf("Error(node_access): etcd create lease has error, %v ", err)
		return errors.New(msg)
	}
	log.Println("lease ok")
	_, err = cli.Put(context.Background(), key, string(value), clientv3.WithLease(lease.ID))
	if err != nil {
		log.Println("lease error")
		msg := fmt.Sprintf("Error(node_access): service info register error, %v", err)
		return errors.New(msg)
	}
	// cli.Put(context.Background(), key, string(value))
	log.Println("put ok")
	return nil
}
