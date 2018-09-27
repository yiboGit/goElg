package main

import (
	"log"

	"eglass.com/utils"
)

func main() {
	request := utils.InitRequest()
	qy := utils.Query{"appid": "wx8bdfeb60fade3728"}
	body := utils.Body{
		"card_id": "p3BDHwLIKboBd9raftZ9p-fxYlY0",
	}
	result, err := request.JPost("https://t.epeijing.cn/api/v2/wechat/card/get", qy, body)
	if err != nil {
		log.Printf("err %v \n", err)
	}
	if result["errcode"].(float64) > 0 {
		log.Printf("result has error %v \n", result["errcode"].(float64))
	}
	log.Printf("%v", result)
}
