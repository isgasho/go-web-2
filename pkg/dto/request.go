package dto

// Login 登录请求的结构体
type Login struct {
	Username string `json:"username" binding:"required,min=10,max=20"`
	Password string `json:"password"`
}
