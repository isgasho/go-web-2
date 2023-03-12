package model

import "github.com/golang-module/carbon/v2"

// User 用户相关数据结构
type User struct {
	Base         `json:"base"`
	Username     string          `gorm:"index:idx_username,unique;comment:用户名" json:"username"`
	Password     string          `gorm:"comment:用户密码" json:"password"`
	Nickname     string          `gorm:"comment:姓名" json:"nickname"`
	Mobile       string          `gorm:"comment:手机号" json:"mobile"`
	Email        string          `gorm:"comment:电子邮箱" json:"email"`
	Avatar       string          `gorm:"comment:头像地址" json:"avatar"`
	UserNumber   string          `gorm:"comment:工号或者学号" json:"userNumber"`
	Introduction string          `gorm:"comment:个人介绍" json:"introduction"`
	Creator      string          `gorm:"default:system;comment:创建人（默认system）" json:"creator"`
	Status       *uint           `gorm:"type:tinyint(1);default:1;comment:用户状态(0: 禁用, 1: 激活)" json:"status"`
	LastLogin    carbon.DateTime `gorm:"comment:最后登录时间" json:"lastLogin"`
	Locked       *uint           `gorm:"type:tinyint(1);default:0;comment:锁定状态(0: 正常, 1: 锁定)" json:"locked"`
	LockExpire   int64           `gorm:"comment:锁定过期时间" json:"lockExpire"`
	WrongTimes   int             `gorm:"comment:用户登录失败次数" json:"wrongTimes"`
}
