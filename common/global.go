package common

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 全局变量
var (
	Config Configuration      // 配置
	Logger *zap.SugaredLogger // 日志输出
	DB     *gorm.DB           // 数据库连接
	Redis  *redis.Client      // Redis 连接
)

// 时间格式化
const (
	MsecLocalTimeFormat = "2006-01-02 15:04:05.000"
	SecLocalTimeFormat  = "2006-01-02 15:04:05"
	DateLocalTimeFormat = "2006-01-02"
)
