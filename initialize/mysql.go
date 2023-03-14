package initialize

import (
	"fmt"
	"go-web/common"
	"go-web/model"
	"go-web/pkg/zapgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strings"
	"time"
)

// Mysql 初始化 MySQL 连接
func Mysql() {
	// 连接字符串，密码 * 是为了日志中打印出来
	dsnLogStr := fmt.Sprintf("%s:******@tcp(%s:%d)/%s?%s&charset=%s&collation=%s",
		common.Config.Mysql.Username,
		common.Config.Mysql.Host,
		common.Config.Mysql.Port,
		common.Config.Mysql.Database,
		common.Config.Mysql.Query,
		common.Config.Mysql.Charset,
		common.Config.Mysql.Collation,
	)

	// 打印数据库连接
	//log.Println("打开 MySQL 链接：", dsnLogStr)

	// 真正连接串
	dsn := strings.Replace(dsnLogStr, "******", common.Config.Mysql.Password, 1)

	// 打开数据库连接
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 170, // string 类型字段的默认长度
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   common.Config.Mysql.TablePrefix, // 表名前缀
			SingularTable: true,                            // 表名单数
		},
		Logger:                                   zapgorm.New(common.Logger),
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键
		QueryFields:                              true, // 解决查询全部字段可能不走索引的问题
	})

	// 连接错误
	if err != nil {
		log.Fatal("数据库连接异常：", err.Error())
	}

	// 设置数据库连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(common.Config.Mysql.MaxIdleConnections)                          // 最大空闲连接数
	sqlDB.SetMaxOpenConns(common.Config.Mysql.MaxOpenConnections)                          // 最大连接数
	sqlDB.SetConnMaxIdleTime(time.Duration(common.Config.Mysql.MaxIdleTime) * time.Minute) // 最大连接复用时间

	// 设置全局
	common.DB = db

	// 日志输出
	log.Println("MySQL 数据库连接成功！")
}

// AutoMigrate 表同步
func AutoMigrate() {
	err := common.DB.AutoMigrate(
		new(model.User),
		new(model.Role),
		new(model.Menu),
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("数据库表结构同步失败：", err.Error()))
	}
	log.Println("数据库表结构同步完成！")
	os.Exit(0)
}
