package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/common"
	"time"
)

// AccessLog 访问日志中间件
func AccessLog(ctx *gin.Context) {
	// 请求时间
	requestTime := time.Now()
	// 处理请求
	ctx.Next()
	// 结束时间
	requestEndTime := time.Now()
	// 请求耗时
	requestExecTime := requestEndTime.Sub(requestTime)
	// 请求方式
	requestMethod := ctx.Request.Method
	// 请求路由
	requestUri := ctx.Request.RequestURI
	// 状态码
	requestCode := ctx.Writer.Status()
	// 请求IP
	requestIP := ctx.ClientIP()
	// 拼接日志信息
	logStr := fmt.Sprintf("%s\t%s\t%d\t%s\t%s", requestMethod, requestUri, requestCode, requestExecTime.String(), requestIP)

	// 判断请求，OPTIONS 请求使用 DEBUG
	if requestMethod == "OPTIONS" {
		common.Logger.Debug(logStr)
	} else {
		common.Logger.Info(logStr)
	}
}
