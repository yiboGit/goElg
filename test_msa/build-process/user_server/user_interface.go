package main

type UserServer interface {
	AddUser(user UserReq) (UserResp, error)
	DeleteUser(user UserReq) (UserResp, error)
	UpdateUser(user UserReq) (UserResp, error)
	GetUser(user UserReq) (UserResp, error)
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

type UserReq struct {
	user  User                   // user  请求参数主体
	extra map[string]interface{} // extra 额外参数
}

type UserResp struct {
	user User // user  返回值主体
}
