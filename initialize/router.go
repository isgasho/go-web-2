package initialize

import (
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

	// api 前缀

	// 创建默认路由组
	bg := r.Group(common.ApiPrefix)
	{
		routes.Public(bg, auth) // 开放路由组

	}

	// 涉及认证鉴权
	ag := r.Group(common.ApiPrefix)
	ag.Use(auth.MiddlewareFunc()) // 认证中间件
	ag.Use(middleware.Casbin)     // 鉴权中间件
	{
		routes.User(ag) // 用户路由组
		routes.Menu(ag) // 菜单路由组
	}

	// 返回路由引擎
	return r
}
