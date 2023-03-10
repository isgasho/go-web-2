package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/common"
	"go-web/pkg/response"
	"runtime/debug"
)

// Exception 异常捕获中间件
func Exception(ctx *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			// 使用断言判断是否为用户响应专门抛出的异常
			resp, ok := err.(response.ResponseInfo)
			if ok {
				// 如果是响应数据抛出的异常
				response.JSON(ctx, response.OK, resp)
				ctx.Abort()
				return
			}

			// 如果是系统本身的异常
			common.Logger.Error(fmt.Sprintf("系统发生未知错误：%s\n详细错误信息：%v", err, string(debug.Stack())))

			// 生成异常响应结构体
			resp = response.ResponseInfo{
				Code:    response.InternalServerError,
				Status:  false,
				Message: response.CustomMessage[response.InternalServerError],
				Data:    map[string]interface{}{},
			}

			response.JSON(ctx, response.OK, resp)
			ctx.Abort()
			return
		}
	}()
	ctx.Next()
}
