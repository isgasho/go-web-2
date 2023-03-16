package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/common"
	"go-web/middleware"
	"go-web/routes"
	"log"
)

func Router() *gin.Engine {
	// gin 基础基础配置
	gin.SetMode(common.Config.System.Mode) // 运行模式
	gin.DisableConsoleColor()              // 关闭控制台颜色输出

	// 创建不带中间件的路由
	r := gin.New()

	// 访问日志中间件
	r.Use(middleware.AccessLog)
	// 跨域中间件
	r.Use(middleware.Cors)
	// 异常捕获中间件
	r.Use(middleware.Exception)

	// JWT 中间件
	auth, err := middleware.JWTAuth()
	if err != nil {
		log.Fatalln("JWT中间件初始化失败：", err.Error())
	}
	// 创建默认路由组
	baseGroup := r.Group(fmt.Sprintf("/%s/%s", common.Config.System.ApiPrefix, common.Config.System.ApiVersion))

	// 开放路由组，不需要任何认证鉴权
	publicGroup := baseGroup.Group("/")
	routes.Public(publicGroup)

	// 认证路由组，登录登出等接口，需要用到 JWT 中间件
	authGroup := baseGroup.Group("/")
	routes.Auth(authGroup, auth)

	// 用户路由组
	userGroup := baseGroup.Group("/")
	routes.User(userGroup, auth)

	// 其它路由组

	// 返回路由引擎
	return r
}
