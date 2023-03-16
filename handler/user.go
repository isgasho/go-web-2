package handler

import (
	"github.com/gin-gonic/gin"
	"go-web/dto"
	"go-web/pkg/response"
	"go-web/pkg/utools"
	"go-web/service/mysql_service"
)

// GetUserInfoById 根据 ID 获取用户信息
func GetUserInfoById(ctx *gin.Context) {
	// 获取需要查询的用户 ID
	idStr := ctx.Query("id")

	// 需要对 ID 进行类型转换
	var userid uint
	userid = utools.String2Uint(idStr)

	// 转换失败或者类型错误
	if userid == 0 {
		response.FailedWithMessage("传入的用户ID不合法")
		return
	}

	// 都正常了，则查询数据库
	s := mysql_service.New()
	user, err := s.GetUserInfoById(userid)
	if err != nil {
		response.FailedWithMessage(err.Error())
		return
	}

	// 获取到了用户数据，则需要对其进行封装，只返回部分数据
	var resp dto.UserInfoResponse
	utools.Struct2StructByJson(user, &resp)

	// 完善返回数据
	resp.Role = user.Role.Name
	response.SuccessWithData(map[string]interface{}{
		"userInfo": resp,
	})
}
