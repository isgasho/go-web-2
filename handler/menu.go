package handler

import (
	"github.com/gin-gonic/gin"
	"go-web/pkg/response"
	"go-web/service/mysql_service"
)

// GetMenuTree 获取菜单树
func GetMenuTree(ctx *gin.Context) {
	// 查询数据
	s := mysql_service.New()
	tree, err := s.GetMenuTree()
	if err != nil {
		response.FailedWithMessage(err.Error())
		return
	}

	// 返回数据
	response.SuccessWithData(map[string]interface{}{
		"menus": tree,
	})
}
