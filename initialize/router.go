package initialize

import (
	"github.com/gin-gonic/gin"
	"go-web/routes"
)

func Router() *gin.Engine {
	// gin 基础基础配置
	gin.SetMode("debug")      // 运行模式
	gin.DisableConsoleColor() // 关闭控制台颜色输出

	// 创建不带中间件的路由
	r := gin.New()

	// 创建默认路由组
	baseGroup := r.Group("/api/v1")

	// 开放路由组
	publicGroup := baseGroup.Group("/public")
	{
		routes.Public(publicGroup)
	}

	// 其它路由组

	// 返回路由引擎
	return r
}
