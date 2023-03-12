package routes

import (
	"github.com/gin-gonic/gin"
	"go-web/handler"
)

// Public 路由组
func Public(rg *gin.RouterGroup) gin.IRoutes {
	rg.GET("/ping", handler.Ping)
	rg.POST("/login", handler.Login)
	return rg
}
