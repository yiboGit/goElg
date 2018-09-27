package main

import brisk "eglass.com/brisk/msa_rpc"

type UserServer interface {
	AddUser(user UserReq) UserResp
	DeleteUser(user UserReq) UserResp
	UpdateUser(user UserReq) UserResp
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

type AddReq struct {
	User User `json:"user"`
}

type AddResp struct {
	User User `json:"user"`
}

type DeleteReq struct {
	User User `json:"user"`
}

type DeleteResp struct {
	User User `json:"user"`
}

type UpdateReq struct {
	User User `json:"user"`
}

type UpdateResp struct {
	User User `json:"user"`
}

type GetReq struct {
	User User `json:"user"`
}

type GetResp struct {
	User User `json:"user"`
}
