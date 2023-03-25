package mysql_service

import (
	"go-web/model"
	"go-web/pkg/utools"
)

// GenerateMenuTree 生成菜单树
func GenerateMenuTree(parentId uint, roleMenuIds []uint, menus []model.Menu) (tree []model.Menu) {
	tree = make([]model.Menu, 0)
	for _, menu := range menus {
		// 菜单 ID 不在角色的菜单 ID 中，则跳过
		if !utools.ContainsUint(roleMenuIds, menu.Id) {
			continue
		}

		if menu.ParentId == parentId {
			// 递归生成每个菜单和它的子菜单
			menu.Children = GenerateMenuTree(menu.Id, roleMenuIds, menus)
			tree = append(tree, menu)
		}
	}
	return
}

// GetMenuTreeByRoleId 根据角色 ID 获取菜单树
func (s *MysqlService) GetMenuTreeByRoleId(roleId uint) (tree []model.Menu, err error) {
	// 初始化
	tree = make([]model.Menu, 0)
	menus := make([]model.Menu, 0)

	var role model.Role
	var roleMenuIds []uint

	// 查询角色对应 RoleMenusIds
	err = s.db.Preload("Menus", "status = ?", true).Where("id", roleId).First(&role).Error
	if err != nil {
		return
	}

	// 生成角色的菜单 ID 列表
	for _, menu := range role.Menus {
		roleMenuIds = append(roleMenuIds, menu.Id)
	}

	// 查询所有菜单数据
	err = s.db.Find(&menus).Error
	if err != nil {
		return
	}

	// 生成菜单树，parentId=0 说明是最外层菜单
	tree = GenerateMenuTree(0, roleMenuIds, menus)
	return tree, nil
}
