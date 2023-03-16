package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go-web/handler"
)

// User 用户路由组
func User(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	// 使用登录之后才能访问的中间件
	rs := rg.Use(auth.MiddlewareFunc())
	{
		rs.GET("/user", handler.GetUserInfoById) // 根据 ID 获取用户信息
	}
	return rs
}
