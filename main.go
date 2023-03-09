package main

import (
	"context"
	"go-web/initialize"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 路由初始化
	router := initialize.Router()

	// 配置服务启动参数
	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	// 启动服务
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("服务异常：", err.Error())
		}
	}()

	// 等待中断信号，实现优雅的关闭
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("服务即将开始关闭...")

	// 等待 5 秒处理，然后停止服务
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal("服务停止异常：", err.Error())
	}
	log.Println("服务关闭完成")
}
