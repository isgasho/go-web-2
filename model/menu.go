package model

// Menu 菜单模型
type Menu struct {
	Base
	Name     string `gorm:"comment:菜单名称" json:"name"`
	Title    string `gorm:"comment:菜单别称" json:"title"`
	Icon     string `gorm:"comment:菜单图标" json:"icon"`
	Path     string `gorm:"comment:菜单路径" json:"path"`
	Status   *uint  `gorm:"type:tinyint(1);default:1;comment:菜单状态(0: 禁用, 1: 激活)" json:"status"`
	Visible  *uint  `gorm:"type:tinyint(1);default:1;comment:显示状态(0: 隐藏, 1: 可见)" json:"visible"`
	ParentId uint   `gorm:"comment:父菜单Id" json:"parentId"`
	Children []Menu `gorm:"-" json:"children"`
	// 菜单角色关系，一个菜单属于多个角色
	Roles []Role `gorm:"many2many:relation_role_menu" json:"roles"`
}
