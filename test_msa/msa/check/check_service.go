package main

import (
	"log"

	"github.com/labstack/echo"
)

var (
	instance = CheckInstace{
		nil,
		echo.New(),
	}
)

type CheckInstace struct {
	err error
	e   *echo.Echo
}

func main() {
	// 生成 User Client端
	// brisk.BuildClient((*UserServer)(nil), "client", "user_service", ".")

	log.Printf("Init")
	// 获取服务实例
	err := uClient.Init()
	if err != nil {
		log.Printf("%v", err)
		return
	}

	// check add function
	log.Printf("测试AddUser-start")
	// reqid

	resp := uClient.AddUser(UserReq{
		User{
			"",
			"001-test",
			"123456",
		},
	})
	log.Printf("resp: %v", resp)
	log.Printf("测试AddUser-end")

	// log.Printf("测试GetUser-start")
	// resp := uClient.GetUser(UserReq{
	// 	User{
	// 		"1111",
	// 		"",
	// 		"",
	// 	},
	// })
	// log.Printf("this resp is : %v", resp)
	// log.Printf("测试GetUser-end")

	// ticker := time.NewTicker(2 * time.Second)

	// for {
	// 	select {
	// 	case <-ticker.C:
	// 		// check add function
	// 		log.Printf("测试AddUser-start")
	// 		resp, err := uClient.AddUser(UserReq{
	// 			User{
	// 				"",
	// 				"001-test",
	// 				"123456",
	// 			},
	// 		})
	// 		if err != nil {
	// 			log.Printf("Err: %v", err)
	// 		}
	// 		log.Printf("resp: %v", resp)
	// 		log.Printf("测试AddUser-end")
	// 	}
	// }

	// // check update function
	// log.Printf("测试UpdateUser-start")
	// client.UpdateUser(UserReq{
	// 	User{
	// 		"",
	// 		"001-test",
	// 		"123456",
	// 	},
	// })
	// log.Printf("测试UpdateUser-end")

	// // check delete function
	// log.Printf("测试DeleteUser-start")
	// client.DeleteUser(UserReq{
	// 	User{
	// 		"",
	// 		"",
	// 		"",
	// 	},
	// })
	// log.Printf("测试DeleteUser-end")

}
