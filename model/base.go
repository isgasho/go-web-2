package model

import (
	"github.com/golang-module/carbon/v2"
)

// Base 基础结构体
type Base struct {
	Id        uint            `gorm:"primaryKey;comment:自增编号" json:"id"`
	CreatedAt carbon.DateTime `gorm:"comment:创建时间" json:"createdAt"`
	UpdatedAt carbon.DateTime `gorm:"comment:更新时间" json:"updatedAt"`
	DeletedAt DeletedAt       `gorm:"index:idx_deleted_at;comment:软删除时间" json:"deletedAt"`
}
