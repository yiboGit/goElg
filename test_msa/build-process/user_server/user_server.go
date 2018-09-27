package main

import (
	"fmt"
	"log"

	"github.com/rs/xid"

	"github.com/labstack/echo"
)

var (
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

func (u UserInstace) AddUser(userReq UserReq) (UserResp, error) {
	log.Println("进入方法")
	gid := xid.New()
	user := userReq.user
	user.ID = fmt.Sprintf("%s", gid)
	cache[user.ID] = user
	log.Printf("添加完成-ID:  %s , %v \n", user.ID, user)
	return UserResp{}, nil
}

func (u UserInstace) DeleteUser(userReq UserReq) (UserResp, error) {
	user := userReq.user
	if cache == nil {
		log.Println("无当前User" + user.ID)
	}
	delete(cache, user.ID)
	log.Printf("删除完成-ID: %s \n", user.ID)
	return UserResp{}, nil
}

func (u UserInstace) GetUser(userReq UserReq) (UserResp, error) {
	user := userReq.user
	if cache == nil {
		log.Println("无当前User" + user.ID)
	}
	resultUser := cache[user.ID]
	return UserResp{resultUser}, nil
}

func (u UserInstace) UpdateUser(userReq UserReq) (UserResp, error) {
	log.Println("进入方法")
	user := userReq.user
	if cache == nil {
		log.Println("无目标User" + user.ID)
	} else if _, ok := cache[user.ID]; ok {
		log.Println("无目标User" + user.ID)
	}
	cache[user.ID] = user
	log.Printf("更新完成-ID: %s , %v \n", user.ID, user)
	return UserResp{}, nil
}

func main() {
	e := u.e
	u.bind()
	e.Logger.Fatal(e.Start(":8099"))
}
