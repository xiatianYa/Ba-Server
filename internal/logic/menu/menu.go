package menu

import (
	v1 "Ba-Server/api/menu/v1"
	"Ba-Server/internal/dao"
	"Ba-Server/internal/model/entity"
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"
	"fmt"
)

type (
	sMenu struct{}
)

func init() {
	service.RegisterMenu(New())
}

func New() service.IMenu {
	return &sMenu{}
}

func (s sMenu) GetAllSysMenus(ctx context.Context, req *v1.GetAllSysMenusReq) (res []string, err error) {
	var Menus []string
	var SysMenus []entity.SysMenu
	menuModel := dao.SysMenu.Ctx(ctx)
	_ = menuModel.Scan(&SysMenus)
	for _, menu := range SysMenus {
		Menus = append(Menus, menu.RouteName)
	}
	return Menus, nil
}

func (s sMenu) GetSysMenuTree(ctx context.Context, req *v1.GetSysMenuTreeReq) (res []vo.SysMenuTreeVo, err error) {
	var SysMenus []entity.SysMenu
	menuModel := dao.SysMenu.Ctx(ctx)
	_ = menuModel.Scan(&SysMenus)
	tree := buildTree(SysMenus)
	return tree, nil
}

func (s sMenu) GetMenuIdsByRoleId(ctx context.Context, req *v1.GetMenuIdsByRoleIdReq) (res []int64, err error) {
	var SysRoleMenus []entity.SysRoleMenu
	var SysMenuIds []int64
	roleMenuModel := dao.SysRoleMenu.Ctx(ctx)
	_ = roleMenuModel.Where("role_id", req.Id).Scan(&SysRoleMenus)
	for _, roleMenu := range SysRoleMenus {
		SysMenuIds = append(SysMenuIds, roleMenu.Id)
	}
	return SysMenuIds, nil
}

// buildTree 构建菜单树
func buildTree(menus []entity.SysMenu) []vo.SysMenuTreeVo {
	// 首先创建一个映射，用于快速查找每个菜单的子菜单
	menuMap := make(map[int64][]vo.SysMenuTreeVo)
	for _, menu := range menus {
		menuTreeVo := vo.SysMenuTreeVo{
			Id:    menu.Id,
			Label: menu.MenuName,
			Pid:   fmt.Sprintf("%d", menu.ParentId),
			Sort:  fmt.Sprintf("%d", menu.Order),
		}
		menuMap[menu.ParentId] = append(menuMap[menu.ParentId], menuTreeVo)
	}

	// 然后构建根菜单列表
	var rootMenus []vo.SysMenuTreeVo
	for _, menu := range menus {
		if menu.ParentId == 0 {
			rootMenu := vo.SysMenuTreeVo{
				Id:    menu.Id,
				Label: menu.MenuName,
				Pid:   fmt.Sprintf("%d", menu.ParentId),
				Sort:  fmt.Sprintf("%d", menu.Order),
			}
			rootMenu.Children = buildChildren(menu.Id, menuMap)
			rootMenus = append(rootMenus, rootMenu)
		}
	}
	return rootMenus
}

// buildChildren 递归构建子菜单
func buildChildren(parentID int64, menuMap map[int64][]vo.SysMenuTreeVo) []vo.SysMenuTreeVo {
	var children []vo.SysMenuTreeVo
	if childList, ok := menuMap[parentID]; ok {
		for _, child := range childList {
			child.Children = buildChildren(child.Id, menuMap)
			children = append(children, child)
		}
	}
	return children
}
