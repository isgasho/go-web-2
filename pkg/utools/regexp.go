package utools

import (
	"regexp"
)

// UsernameReg 常用正则
const (
	UsernameReg = `^[a-z][a-z0-9]{3,20}$`                                               // 用户名正则
	EmailReg    = `^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$` // 邮箱正则
	MobileReg   = `^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\d{8}$`          // 手机号正则
)

// RegexpString 正则校验
func RegexpString(reg string, str string) bool {
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(str)
}
