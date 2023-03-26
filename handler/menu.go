package handler

import (
	"github.com/gin-gonic/gin"
	"go-web/pkg/response"
	"go-web/pkg/utools"
	"go-web/service/mysql_service"
)

// GetMenuTreeByRoleId 获取角色获取菜单树
func GetMenuTreeByRoleId(ctx *gin.Context, roleId uint) {
	// 查询数据
	s := mysql_service.New()
	tree, err := s.GetMenuTreeByRoleId(roleId)
	if err != nil {
		response.FailedWithMessage(err.Error())
		return
	}

	// 返回数据
	response.SuccessWithData(map[string]interface{}{
		"menus": tree,
	})
}

// GetRoleMenuTree 根据传递角色 Id 查询菜单树
func GetRoleMenuTree(ctx *gin.Context) {
	// 获取传递过来的角色 Id
	idStr := ctx.Param("roleId")
	var roleId uint
	roleId = utools.String2Uint(idStr)

	// 转换失败或者类型错误
	if roleId == 0 {
		response.FailedWithMessage("传入的角色ID不合法")
		return
	}

	// 根据角色 Id 查询菜单树
	GetMenuTreeByRoleId(ctx, roleId)
}

// GetCurrentUserMenuTree 获取当前用户的菜单树
func GetCurrentUserMenuTree(ctx *gin.Context) {
	// 获取当前用户信息
	user, err := GetCurrentUserFromContext(ctx)
	if err != nil {
		response.FailedWithMessage("获取当前用户的角色 Id 失败")
		return
	}

	// 根据角色 Id 查询菜单树
	GetMenuTreeByRoleId(ctx, user.Role.Id)
}
