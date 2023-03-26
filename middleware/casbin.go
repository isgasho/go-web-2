package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web/common"
	"go-web/handler"
	"go-web/pkg/response"
	"strings"
	"sync"
)

var checkLock sync.Mutex

// Casbin 中间件
func Casbin(ctx *gin.Context) {
	// 获取当前用户信息
	user, _ := handler.GetCurrentUserFromContext(ctx)
	// sub，角色关键字
	sub := user.Role.Keyword
	// obj，除去前缀后的 URI
	apiPrefix := "/" + common.Config.System.ApiPrefix + "/" + common.Config.System.ApiVersion
	obj := strings.TrimPrefix(ctx.Request.RequestURI, apiPrefix)
	// act，请求方式
	act := ctx.Request.Method

	// 设置同一时间只允许一个校验，否则可能出现校验失败
	checkLock.Lock()
	defer checkLock.Unlock()

	// 校验数据
	pass, _ := common.CasbinEnforcer.Enforce(sub, obj, act)
	if !pass {
		response.FailedWithCode(response.Forbidden)
		ctx.Abort()
		return
	}
	ctx.Next()
}
