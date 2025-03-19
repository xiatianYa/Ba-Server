package menu

import (
	v1 "Ba-Server/api/menu/v1"
	"Ba-Server/internal/dao"
	"Ba-Server/internal/model/domain"
	"Ba-Server/internal/model/entity"
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
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

func (s sMenu) GetSysMenuIdsByRoleId(ctx context.Context, req *v1.GetSysMenuIdsByRoleIdReq) (res []int64, err error) {
	var SysRoleMenus []entity.SysRoleMenu
	var SysMenuIds []int64
	roleMenuModel := dao.SysRoleMenu.Ctx(ctx)
	_ = roleMenuModel.Where("role_id", req.Id).Scan(&SysRoleMenus)
	for _, roleMenu := range SysRoleMenus {
		SysMenuIds = append(SysMenuIds, roleMenu.MenuId)
	}
	return SysMenuIds, nil
}

func (s sMenu) GetSysMenuPage(ctx context.Context, req *v1.GetSysMenuPageReq) (total int, records []vo.SysMenuVo, err error) {
	// 创建指针切片
	var SysMenus []entity.SysMenu

	menuModel := dao.SysMenu.Ctx(ctx)
	pageQuery := menuModel.Page(req.Current, req.Size)

	err = pageQuery.ScanAndCount(&SysMenus, &total, true)

	if err != nil {
		return 0, nil, err
	}

	//对数据进行处理
	tree := buildVoTree(SysMenus, ctx)

	records = tree
	return
}

func (s sMenu) SaveSysMenu(ctx context.Context, req *v1.SaveSysMenuReq) (res *v1.SaveSysMenuRes, err error) {
	return
}

// buildVoTree 构建菜单Vo树
func buildVoTree(menus []entity.SysMenu, ctx context.Context) []vo.SysMenuVo {
	// 首先创建一个映射，用于快速查找每个菜单的子菜单
	menuMap := make(map[int64][]vo.SysMenuVo)
	for _, menu := range menus {
		var buttons []domain.BTPairs
		var sysPermissions []entity.SysPermission
		_ = dao.SysPermission.Ctx(ctx).Where("menu_id", menu.Id).Scan(&sysPermissions)
		for _, permission := range sysPermissions {
			button := domain.BTPairs{
				Code: permission.Code,
				Desc: permission.Description,
			}
			buttons = append(buttons, button)
		}
		var query []domain.KVPairs
		_ = gconv.Struct(menu.Query, &query)

		//查询菜单的按钮列表
		menuTreeVo := vo.SysMenuVo{
			Id:              menu.Id,
			ParentId:        menu.ParentId,
			MenuType:        menu.MenuType,
			MenuName:        menu.MenuName,
			I18nKey:         menu.I18NKey,
			RouteName:       menu.RouteName,
			RoutePath:       menu.RoutePath,
			Icon:            menu.Icon,
			IconType:        menu.IconType,
			Component:       menu.Component,
			KeepAlive:       strToBoolPtrIgnoreCase(menu.KeepAlive),
			HideInMenu:      strToBoolPtrIgnoreCase(menu.HideInMenu),
			Constant:        strToBoolPtrIgnoreCase(menu.Constant),
			Href:            menu.Href,
			ActiveMenu:      menu.ActiveMenu,
			Order:           menu.Order,
			MultiTab:        strToBoolPtrIgnoreCase(menu.MultiTab),
			FixedIndexInTab: menu.FixedIndexInTab,
			Query:           query,
			Buttons:         buttons,
			Status:          menu.Status,
		}
		menuMap[menu.ParentId] = append(menuMap[menu.ParentId], menuTreeVo)
	}

	// 然后构建根菜单列表
	var rootMenus []vo.SysMenuVo
	for _, menu := range menus {
		if menu.ParentId == 0 {
			var buttons []domain.BTPairs
			var sysPermissions []entity.SysPermission
			_ = dao.SysPermission.Ctx(ctx).Where("menu_id", menu.Id).Scan(&sysPermissions)
			for _, permission := range sysPermissions {
				button := domain.BTPairs{
					Code: permission.Code,
					Desc: permission.Description,
				}
				buttons = append(buttons, button)
			}
			var query []domain.KVPairs
			_ = gconv.Struct(menu.Query, &query)
			rootMenu := vo.SysMenuVo{
				Id:              menu.Id,
				ParentId:        menu.ParentId,
				MenuType:        menu.MenuType,
				MenuName:        menu.MenuName,
				I18nKey:         menu.I18NKey,
				RouteName:       menu.RouteName,
				RoutePath:       menu.RoutePath,
				Icon:            menu.Icon,
				IconType:        menu.IconType,
				Component:       menu.Component,
				KeepAlive:       strToBoolPtrIgnoreCase(menu.KeepAlive),
				HideInMenu:      strToBoolPtrIgnoreCase(menu.HideInMenu),
				Constant:        strToBoolPtrIgnoreCase(menu.Constant),
				Href:            menu.Href,
				ActiveMenu:      menu.ActiveMenu,
				Order:           menu.Order,
				MultiTab:        strToBoolPtrIgnoreCase(menu.MultiTab),
				FixedIndexInTab: menu.FixedIndexInTab,
				Query:           query,
				Buttons:         buttons,
				Status:          menu.Status,
			}
			rootMenu.Children = buildVoChildren(menu.Id, menuMap)
			rootMenus = append(rootMenus, rootMenu)
		}
	}
	return rootMenus
}

// IsYIgnoreCase 函数用于忽略大小写判断字符串是否为 "Y"
func strToBoolPtrIgnoreCase(s string) *bool {
	b := strings.EqualFold(s, "Y") // 忽略大小写比较
	return &b
}

// buildVoChildren 递归构建子菜单
func buildVoChildren(parentID int64, menuMap map[int64][]vo.SysMenuVo) []vo.SysMenuVo {
	var children []vo.SysMenuVo
	if childList, ok := menuMap[parentID]; ok {
		for _, child := range childList {
			child.Children = buildVoChildren(child.Id, menuMap)
			children = append(children, child)
		}
	}
	return children
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
