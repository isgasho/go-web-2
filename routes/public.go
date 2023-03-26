package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go-web/handler"
)

// Public 开放路由组
func Public(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/ping", handler.Ping)        // 状态测试
	rg.POST("/login", auth.LoginHandler) // 用户登录
	return rg
}
