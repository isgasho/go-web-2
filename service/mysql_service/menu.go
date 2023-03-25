package mysql_service

import (
	"go-web/model"
)

// GenerateMenuTree 生成菜单树
func GenerateMenuTree(parentId uint, menus []model.Menu) (tree []model.Menu) {
	tree = make([]model.Menu, 0)
	for _, menu := range menus {
		if menu.ParentId == parentId {
			// 递归生成每个菜单和它的子菜单
			menu.Children = GenerateMenuTree(menu.Id, menus)
			tree = append(tree, menu)
		}
	}
	return
}

// GetMenuTree 获取菜单树
func (s *MysqlService) GetMenuTree() (tree []model.Menu, err error) {
	tree = make([]model.Menu, 0)
	menus := make([]model.Menu, 0)

	// 查询所有菜单数据
	err = s.db.Find(&menus).Error
	if err != nil {
		return
	}

	// 生成菜单树，parentId=0 说明是最外层菜单
	tree = GenerateMenuTree(0, menus)
	return tree, nil
}
