package dto

import (
	"github.com/golang-module/carbon/v2"
)

// UserInfoResponse 用户信息响应结构体
type UserInfoResponse struct {
	Id           uint            `json:"id"`
	Username     string          `json:"username"`
	Nickname     string          `json:"nickname"`
	Mobile       string          `json:"mobile"`
	Email        string          `json:"email"`
	Avatar       string          `json:"avatar"`
	UserNumber   string          `json:"userNumber"`
	Introduction string          `json:"introduction"`
	Creator      string          `json:"creator"`
	Status       *uint           `json:"status"`
	LastLogin    carbon.DateTime `json:"lastLogin"`
	Locked       *uint           `json:"locked"`
	Role         struct {
		Id      uint   `json:"id"`
		Name    string `json:"name"`
		Keyword string `json:"keyword"`
		Desc    string `json:"desc"`
	} `json:"role"`
	CreatedAt carbon.DateTime `json:"createdAt"`
	UpdatedAt carbon.DateTime `json:"updatedAt"`
}
