package handler

import (
	"github.com/gin-gonic/gin"
	"go-web/pkg/response"
	"net/http"
)

// Ping 接口处理函数
func Ping(c *gin.Context) {
	response.JSON(c, http.StatusOK, response.SuccessWithData(map[string]interface{}{
		"info": "pong",
	}))
}
