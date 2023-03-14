package initialize

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-web/common"
	"log"
	"time"
)

// Redis 初始化 Redis
func Redis() {
	dsn := fmt.Sprintf("%s:%d", common.Config.Redis.Host, common.Config.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Network:         "tcp",                        // 协议
		Addr:            dsn,                          // 连接地址
		Password:        common.Config.Redis.Password, // 密码
		DB:              common.Config.Redis.DB,       // 数据库
		MaxRetries:      0,                            // 执行失败重试次数
		MinRetryBackoff: 8 * time.Millisecond,         // 每次计算重试的间隔下限
		MaxRetryBackoff: 512 * time.Millisecond,       // 每次计算重试的间隔上限
		DialTimeout:     5 * time.Second,              // 连接建立超时时间
		ReadTimeout:     3 * time.Second,              // 读取数据超时时间
		WriteTimeout:    3 * time.Second,              // 写超时时间
		PoolSize:        14,                           // 连接池最大连接数，一般和 CPU 核数 4 倍少一点
		PoolTimeout:     4 * time.Second,              // 连接池占满，客户端最大等待时常，读 + 1 秒
		MinIdleConns:    10,                           // 在启动的时候默认创建的连接数量
		ConnMaxIdleTime: 5 * time.Minute,              // 连接最大空闲时间
		ConnMaxLifetime: 0,                            // 连接存活时长
	})

	// 测试连接是否正常
	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalln("Redis 连接失败：", err.Error())
	}
	common.Redis = client
	log.Println("Redis 缓存连接成功！")
}
