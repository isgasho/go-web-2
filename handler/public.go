package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping 接口处理函数
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
