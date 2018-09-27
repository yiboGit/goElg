package main
import (
	"github.com/coreos/etcd/clientv3"
	"time"
	"context"
	"fmt"
)

func m() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close()
	cli.Put(context.Background(), "foo",  "value", clientv3.WithPrefix())
	fooChan := cli.Watch(context.Background(), "foo")
	select {
		case e:= <- fooChan: 
		for _, ev := range e.Events {
			fmt.Printf("type: %s, %s, %s", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
		// fmt.Println(e.Events)
	}
}