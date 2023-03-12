package model

// Role 角色模型
type Role struct {
	Base
	Name    string `gorm:"index:idx_name,unique;comment:角色名称" json:"name"`
	Keyword string `gorm:"index:idx_keyword,unique;comment:角色关键字" json:"keyword"`
	Desc    string `gorm:"comment:角色说明" json:"desc"`
	Status  *uint  `gorm:"type:tinyint(1);default:1;comment:角色状态(0: 禁用, 1: 启用)" json:"status"`
	// 角色外键，用于反查
	Users []User `gorm:"foreignKey:RoleId" json:"users"`
	// 菜单角色关系，一个角色包含多个菜单
	Menus []Menu `gorm:"many2many:relation_role_menu" json:"menus"`
}
