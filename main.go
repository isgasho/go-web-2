package main

import (
	"context"
	"embed"
	"fmt"
	"go-web/common"
	"go-web/initialize"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

/*
embed 用法参考：https://blog.csdn.net/wan212000/article/details/127264475
这里指定 embed 读取 config 目录下的所有文件
*/

//go:embed config/*
var fs embed.FS

func main() {
	// 配置初始化
	initialize.Config(fs)
	// 日志初始化
	initialize.Logger()
	// 初始化数据库连接
	initialize.Mysql()
	// 路由初始化
	router := initialize.Router()

	// 配置服务启动参数
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", common.Config.System.Port),
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
	common.Logger.Warnln("服务即将开始关闭...")

	// 等待 5 秒处理，然后停止服务
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		common.Logger.Errorln("服务停止异常：", err.Error())
	}
	common.Logger.Info("服务关闭完成！")
}
