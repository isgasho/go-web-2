package routes

import (
	"github.com/gin-gonic/gin"
	"go-web/handler"
)

// Menu 菜单由组
func Menu(rg *gin.RouterGroup) gin.IRoutes {
	rs := rg.Group("/menu")
	{
		rs.GET("/tree", handler.GetCurrentUserMenuTree)  // 获取当期用户的菜单树
		rs.GET("/tree/:roleId", handler.GetRoleMenuTree) // 获取角色 Id 的查询菜单树
	}
	return rg
}
