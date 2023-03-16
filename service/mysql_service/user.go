package mysql_service

import (
	"errors"
	"go-web/model"
	"go-web/pkg/response"
	"go-web/pkg/utools"
	"gorm.io/gorm"
)

// LoginCheck 登录校验，支持用户名，邮箱，手机号登录
func (s *MysqlService) LoginCheck(username string, password string) (*model.User, error) {
	// 查询用户信息
	var user model.User

	// 传递过来的 username 可能是用户名，邮箱，手机号，需要进行判断
	var err error
	if utools.RegexpString(utools.UsernameReg, username) {
		err = s.db.Where("username = ?", username).First(&user).Error
	} else if utools.RegexpString(utools.EmailReg, username) {
		err = s.db.Where("email = ?", username).First(&user).Error
	} else if utools.RegexpString(utools.MobileReg, username) {
		err = s.db.Where("mobile = ?", username).First(&user).Error
	} else {
		return nil, errors.New(response.UserLoginErrorMessage)
	}

	// 判断用户是否存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New(response.UserLoginErrorMessage)
	} else {
		// 验证密码
		if !utools.ComparePassword(user.Password, password) {
			return nil, errors.New(response.UserLoginErrorMessage)
		}

		// 判断用户是否被禁用
		if *user.Status == 0 {
			return nil, errors.New(response.UserDisableMessage)
		}

		// 判断用户是否被锁定
		if *user.Locked == 1 {
			return nil, errors.New(response.UserLockedMessage)
		}
	}
	return &user, nil
}
