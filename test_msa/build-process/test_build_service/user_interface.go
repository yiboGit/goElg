package main

type UserServer interface {
	AddUser(user UserReq) (UserResp, error)
	DeleteUser(user UserReq) (UserResp, error)
	UpdateUser(user UserReq) (UserResp, error)
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

type UserReq struct {
	user  User
	extra map[string]interface{}
}

type UserResp struct {
	user User
}
