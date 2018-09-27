package main

import (
	"encoding/json"
	"errors"
	"log"

	brisk "eglass.com/brisk/msa_rpc"
)

// 临时入口类 帮助生成 bind/client 文件
func main() {
	// brisk.BuildService((*UserServer)(nil), "UserInstace", "instance", ".")
	brisk.BuildClient((*UserServer)(nil), "userClient", "user_service", ".")

	err := errors.New("123qwwe错误")
	e, _ := json.Marshal(err)
	log.Printf("%v", e)
	var er error
	json.Unmarshal(e, &er)
	log.Printf("%v", er)
}
