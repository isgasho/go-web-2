package handler

import (
	"github.com/gin-gonic/gin"
	"go-web/common"
	"go-web/model"
	"go-web/pkg/dto"
	"go-web/pkg/response"
)

// Ping 接口处理函数
func Ping(ctx *gin.Context) {
	response.SuccessWithData(map[string]interface{}{
		"info": "pong",
	})
}

// Login 接口
func Login(ctx *gin.Context) {
	var login dto.Login
	var user model.User

	// 获取用户传递的数据
	err := ctx.ShouldBind(&login)
	if err == nil {
		// 判断是的和数据库匹配
		result := common.DB.Where("username = ? and password = ?", login.Username, login.Password).First(&user)
		if result.RowsAffected == 1 {
			// 返回登录成功
			response.SuccessWithData(map[string]interface{}{
				"username": login.Username,
			})
		}
		return
	}

	response.FailedWithCode(response.UserLoginError)
	return
}
