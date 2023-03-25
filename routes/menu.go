package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go-web/handler"
)

// Menu 菜单由组
func Menu(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	// 使用登录之后才能访问的中间件
	rs := rg.Use(auth.MiddlewareFunc())
	{
		rs.GET("/menus/tree", handler.GetMenuTree) // 获取菜单树
	}
	return rs
}
