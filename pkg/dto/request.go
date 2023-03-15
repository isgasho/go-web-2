package dto

// Login 登录请求的结构体
type Login struct {
	Username string `json:"username" binding:"is-username"`
	Password string `json:"password" binding:"min=8"`
}

// FieldTrans 用户登录结构体字段翻译（必须）
func (l Login) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["Username"] = "用户名"
	m["Password"] = "密码"
	return m
}

// FieldErrors 用户登录结构体错误信息（必须，用于自定义规则校验不通过报错）
func (l Login) FieldErrors() map[string]string {
	m := make(map[string]string, 0)
	m["Username"] = "用户名必须为小写字母开头长度为4-30的字符串（只支持字母数字）"
	return m
}
