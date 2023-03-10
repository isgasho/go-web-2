package common

import "go.uber.org/zap"

// 全局变量
var (
	Config Configuration      // 配置
	Logger *zap.SugaredLogger // 日志输出
)

// 时间格式化
const (
	MsecLocalTimeFormat = "2006-01-02 15:04:05.000"
	SecLocalTimeFormat  = "2006-01-02 15:04:05"
	DateLocalTimeFormat = "2006-01-02"
)