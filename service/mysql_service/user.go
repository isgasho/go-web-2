package mysql_service

import (
	"errors"
	"go-web/model"
	"go-web/pkg/response"
	"go-web/pkg/utools"
	"gorm.io/gorm"
)

// LoginByUsername 登录校验
func (s *MysqlService) LoginByUsername(username string, password string) (*model.User, error) {
	// 查询用户信息
	var user model.User
	err := s.db.Where("username = ?", username).First(&user).Error

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
