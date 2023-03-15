package initialize

import (
	"errors"
	"go-web/common"
	"go-web/model"
	"go-web/pkg/utools"
	"gorm.io/gorm"
	"os"
)

// 用户数据
var users = []model.User{
	model.User{
		Username:   "admin",
		Password:   "12345678",
		Nickname:   "超级管理员",
		Mobile:     "18888888888",
		Email:      "abc@qq.com",
		Avatar:     "1.png",
		UserNumber: "wx001",
	},
}

// User 初始化用户数据
func User() {
	for _, user := range users {
		// 查看记录是否存在
		err := common.DB.Where("username = ?", user.Username).First(&user).Error
		// 如果记录不存在则添加数据
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 密码加密存储
			user.Password = utools.CryptoPassword(user.Password)
			common.DB.Create(&user)
		}
	}
	os.Exit(0)
}
