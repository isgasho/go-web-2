package routes

import (
	"github.com/gin-gonic/gin"
	"go-web/handler"
)

// User 用户路由组
func User(rg *gin.RouterGroup) gin.IRoutes {
	rs := rg.Group("/user")
	{
		rs.GET("/", handler.GetUserInfoById) // 根据 ID 获取用户信息
	}
	return rg
}
