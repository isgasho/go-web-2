package dto

// Login 登录请求的结构体
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
