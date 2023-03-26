package initialize

import (
	"errors"
	"go-web/common"
	"go-web/model"
	"go-web/pkg/utools"
	"gorm.io/gorm"
	"os"
)

// 用户数据
var users = []model.User{
	{
		Username:   "admin",
		Password:   "12345678",
		Nickname:   "超级管理员",
		Mobile:     "18888888888",
		Email:      "abc@qq.com",
		Avatar:     "1.png",
		UserNumber: "wx001",
		RoleId:     1,
	},
	{
		Username:   "guest",
		Password:   "12345678",
		Nickname:   "访客",
		Mobile:     "19999999999",
		Email:      "xyz@qq.com",
		Avatar:     "2.png",
		UserNumber: "wx002",
		RoleId:     2,
	},
}

// User 初始化用户数据
func Users() {
	for _, user := range users {
		// 查看记录是否存在
		err := common.DB.Where("username = ?", user.Username).First(&user).Error
		// 如果记录不存在则添加数据
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 密码加密存储
			user.Password = utools.CryptoPassword(user.Password)
			common.DB.Create(&user)
		}
	}
}

// 角色数据
var roles = []model.Role{
	{
		Base: model.Base{
			Id: 1,
		},
		Name:    "超管",
		Keyword: "super",
		Desc:    "超级管理员",
	},
	{
		Base: model.Base{
			Id: 2,
		},
		Name:    "访客",
		Keyword: "guest",
		Desc:    "访客，部分只读权限",
	},
}

// Roles 初始化角色数据
func Roles() {
	for _, role := range roles {
		// 查看记录是否存在
		err := common.DB.Where("id = ? or name = ?", role.Id, role.Name).First(&role).Error
		// 如果记录不存在则添加数据
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.DB.Create(&role)
		}
	}
}

// 菜单数据
var menus = []model.Menu{
	{
		Base: model.Base{
			Id: 1,
		},
		Name:     "工作台",
		Title:    "Dashboard",
		Icon:     "line-chart-outlined",
		Path:     "/dashboard",
		ParentId: 0,
		Roles:    roles,
	},
	{
		Base: model.Base{
			Id: 10,
		},
		Name:     "系统设置",
		Title:    "Setting",
		Icon:     "setting-outlined",
		Path:     "/setting",
		ParentId: 0,
		Roles: []model.Role{
			roles[0],
		},
	},
	{
		Base: model.Base{
			Id: 11,
		},
		Name:     "用户管理",
		Title:    "User",
		Icon:     "user-outlined",
		Path:     "/setting/user",
		ParentId: 10,
		Roles: []model.Role{
			roles[0],
		},
	},
	{
		Base: model.Base{
			Id: 12,
		},
		Name:     "角色管理",
		Title:    "Role",
		Icon:     "usergroup-add-outlined",
		Path:     "/setting/role",
		ParentId: 10,
		Roles: []model.Role{
			roles[0],
		},
	},
	{
		Base: model.Base{
			Id: 13,
		},
		Name:     "菜单管理",
		Title:    "Menu",
		Icon:     "cluster-outlined",
		Path:     "/setting/menu",
		ParentId: 10,
		Roles: []model.Role{
			roles[0],
		},
	},
}

// Menus 初始化角色数据
func Menus() {
	for _, menu := range menus {
		// 查看记录是否存在
		err := common.DB.Where("id = ? or name = ?", menu.Id, menu.Name).First(&menu).Error
		// 如果记录不存在则添加数据
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.DB.Create(&menu)
		}
	}
}

// 初始化 RBAC 鉴权
var casbinRules = []model.CasbinRule{
	{
		PType:   "p",
		Keyword: "guest",
		Path:    "/dashboard",
		Method:  "GET",
	},
	{
		PType:   "p",
		Keyword: "guest",
		Path:    "/user",
		Method:  "GET",
	},
	{
		PType:   "p",
		Keyword: "guest",
		Path:    "/menu/tree",
		Method:  "GET",
	},
}

// CasbinRules 初始化鉴权数据
func CasbinRules() {
	for _, rule := range casbinRules {
		// 查看记录是否存在
		err := common.DB.Where(rule).First(&rule).Error
		// 如果记录不存在则添加数据
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.DB.Create(&rule)
		}
	}
}

// Data 初始化数据
func Data() {
	Roles()       // 角色
	Users()       // 用户
	Menus()       // 菜单
	CasbinRules() // 鉴权
	os.Exit(0)
}
