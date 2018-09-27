package main

import (
	"errors"
	"fmt"
	"log"

	"eglass.com/brisk"
	// "eglass.com/brisk/msa_rpc"
	msa "eglass.com/brisk/msa_rpc"
	"github.com/rs/xid"

	"github.com/labstack/echo"
)

var (
	// Port 端口
	Port = "8099"
	// Host 域名
	Host = "localhost"
	// ContainerPort os获取配置文件上容器端口
	ContainerPort = "8099"

	cache = make(map[string]User)
	u     = UserInstace{
		nil,
		echo.New(),
	}
)

type UserInstace struct {
	err error
	e   *echo.Echo
}

func (u UserInstace) AddUser(userReq UserReq) UserResp {
	log.Println("进入方法")
	gid := xid.New()
	user := userReq.User
	user.ID = fmt.Sprintf("%s", gid)
	cache[user.ID] = user
	log.Printf("添加完成-ID:  %s , %v \n", user.ID, user)
	return UserResp{}
}

func (u UserInstace) GetUser(userReq UserReq) UserResp {
	user := userReq.User
	if cache == nil {
		log.Printf("无当前User %v", user.ID)
		return UserResp{
			User{},
			msa.MsError{Message: "无当前User"},
		}
	}
	resultUser, ok := cache[user.ID]
	if !ok {
		err := errors.New("无此数据（user）")
		log.Printf("无当前User %v %v", user.ID, err)
		return UserResp{
			User{},
			msa.MsError{Message: "无当前User"},
		}
	}
	return UserResp{resultUser, msa.MsError{}}
}

// 实现服务注册接口
type ServerInstance struct{}

func (s ServerInstance) Service() error {
	e := u.e
	u.bind()
	return e.Start(":8099")
}

func main() {
	server := ServerInstance{}
	brisk.HandleServiceLifeCycle(server)
}
