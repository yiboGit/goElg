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
	user  User                   // user 请求体主体
	extra map[string]interface{} // extra 请求体额外参数
}

type UserResp struct {
	user User
}
