package dto

// LoginRequest 登录请求的结构体
type LoginRequest struct {
	Username string `json:"username" binding:"min=4,max=30"`
	Password string `json:"password" binding:"min=8"`
}

// FieldTrans 用户登录结构体字段翻译（必须）
func (l LoginRequest) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["Username"] = "用户名"
	m["Password"] = "密码"
	return m
}

// FieldErrors 用户登录结构体错误信息（必须，用于自定义规则校验不通过报错）
func (l LoginRequest) FieldErrors() map[string]string {
	m := make(map[string]string, 0)
	m["Username"] = "请输入合法的用户名，邮箱或者手机号"
	return m
}
