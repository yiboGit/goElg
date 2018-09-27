package main

import brisk "eglass.com/brisk/msa_rpc"

type UserServer interface {
	AddUser(user UserReq) UserResp
	GetUser(user UserReq) UserResp
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

type UserReq struct {
	User User `json:"user"` // user 请求体主体
}

type UserResp struct {
	User User          `json:"user"`
	Err  brisk.MsError `json:"err"`
}
