package mysql_service

import (
	"go-web/common"
	"gorm.io/gorm"
)

// MysqlService 实例对象
type MysqlService struct {
	db *gorm.DB
}

// New 创建实例对象，获得一个 MySQL 连接
func New() MysqlService {
	return MysqlService{db: common.DB}
}
