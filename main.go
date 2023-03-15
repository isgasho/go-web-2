package main

import (
	"context"
	"embed"
	"fmt"
	"go-web/common"
	"go-web/initialize"
	"gorm.io/gorm/utils"
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

// 默认运行环境
var runEnv = "dev"
var command = ""
var commands = []string{"help", "migrate", "init"}

//go:embed config/*
var fs embed.FS

func main() {
	// 判断用户传递的参数
	if len(os.Args) == 2 && utils.Contains(commands, os.Args[1]) {
		command = os.Args[1]
	} else if len(os.Args) == 3 && os.Args[1] == "run" {
		runEnv = os.Args[2]
	} else {
		printHelp()
		os.Exit(1)
	}

	// 配置初始化
	initialize.Config(fs, runEnv)
	// 日志初始化
	initialize.Logger()
	// 初始化数据库连接
	initialize.Mysql()
	// 数据表同步
	if command == "migrate" {
		initialize.AutoMigrate()
	}
	// 数据初始化
	if command == "init" {
		initialize.User()
	}
	// 初始化 Redis 连接
	initialize.Redis()
	// 初始化验证器
	initialize.Validator()
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

// 打印帮助方法
func printHelp() {
	helpInfo := `
使用方法：
	go-web [参数1] [参数2...]

参数说明：
	help：查看帮助信息
	migrate：同步数据结构
	init：初始化用户数据
	run [env]：指定运行环境，默认 dev
`
	fmt.Println(helpInfo)
}
