package handler

import (
	"github.com/gin-gonic/gin"
	"go-web/pkg/response"
)

// Ping 接口处理函数
func Ping(c *gin.Context) {
	response.SuccessWithData(map[string]interface{}{
		"info": "pong",
	})
}
