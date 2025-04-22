package menu

import (
	v1 "Ba-Server/api/menu/v1"
	"Ba-Server/internal/consts"
	"Ba-Server/internal/dao"
	"Ba-Server/internal/model/domain"
	"Ba-Server/internal/model/entity"
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
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
	pageQuery := menuModel.Order("order").Page(req.Current, req.Size)

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
	query, err := gjson.EncodeString(req.Query)
	permissionModel := dao.SysPermission.Ctx(ctx)
	menuModel := dao.SysMenu.Ctx(ctx)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		//查询菜单的按钮列表
		sysMenu := entity.SysMenu{
			ParentId:        req.ParentID,
			MenuType:        req.MenuType,
			MenuName:        req.MenuName,
			I18NKey:         req.I18nKey,
			RouteName:       req.RouteName,
			RoutePath:       req.RoutePath,
			Icon:            req.Icon,
			IconType:        req.IconType,
			Component:       req.Component,
			KeepAlive:       boolToStrPtrIgnoreCase(req.KeepAlive),
			HideInMenu:      boolToStrPtrIgnoreCase(req.HideInMenu),
			Constant:        boolToStrPtrIgnoreCase(req.Constant),
			Href:            req.Href,
			Order:           req.Order,
			MultiTab:        boolToStrPtrIgnoreCase(req.MultiTab),
			FixedIndexInTab: req.FixedIndexInTab,
			Query:           query,
			Status:          req.Status,
		}
		//添加菜单
		menuId, err01 := menuModel.InsertAndGetId(sysMenu)
		if err01 != nil {
			return err01
		}
		//添加菜单按钮
		buttons := req.Buttons
		for _, button := range buttons {
			sysPermission := entity.SysPermission{
				MenuId:      menuId,
				MenuName:    req.MenuName,
				Code:        button.Code,
				Description: button.Desc,
				Status:      consts.ONE,
			}
			_, err02 := permissionModel.Insert(sysPermission)
			if err02 != nil {
				return err02
			}
		}
		return nil
	})
	return
}

func (s sMenu) RemoveSysMenuByIds(ctx context.Context, req *v1.RemoveSysMenuByIdsReq) (res *v1.RemoveSysMenuByIdsRes, err error) {
	var Permissions []entity.SysPermission
	var PermissionIds []int64
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		menuModel := dao.SysMenu.Ctx(ctx)
		permissionModel := dao.SysPermission.Ctx(ctx)
		rolePermissionModel := dao.SysRolePermission.Ctx(ctx)
		roleMenuModel := dao.SysRoleMenu.Ctx(ctx)
		err = permissionModel.WhereIn("menu_id", req.Ids).Scan(&Permissions)
		for _, permission := range Permissions {
			PermissionIds = append(PermissionIds, permission.Id)
		}
		if err != nil {
			return err
		}
		//删除菜单
		_, err = menuModel.WhereIn("id", req.Ids).Delete()
		if err != nil {
			return err
		}
		//删除按钮
		_, err = permissionModel.WhereIn("id", PermissionIds).Delete()
		if err != nil {
			return err
		}
		//删除角色按钮
		_, err = rolePermissionModel.WhereIn("permission_id", PermissionIds).Delete()
		if err != nil {
			return err
		}
		//删除角色菜单
		_, err = roleMenuModel.WhereIn("menu_id", req.Ids).Delete()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}

func (s sMenu) RemoveSysMenuById(ctx context.Context, req *v1.RemoveSysMenuByIdReq) (res *v1.RemoveSysMenuByIdRes, err error) {
	var Permissions []entity.SysPermission
	var PermissionIds []int64
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		menuModel := dao.SysMenu.Ctx(ctx)
		permissionModel := dao.SysPermission.Ctx(ctx)
		rolePermissionModel := dao.SysRolePermission.Ctx(ctx)
		roleMenuModel := dao.SysRoleMenu.Ctx(ctx)
		err = permissionModel.Where("menu_id", req.Id).Scan(&Permissions)
		for _, permission := range Permissions {
			PermissionIds = append(PermissionIds, permission.Id)
		}
		if err != nil {
			return err
		}
		//删除菜单
		_, err = menuModel.Where("id", req.Id).Delete()
		if err != nil {
			return err
		}
		//删除按钮
		_, err = permissionModel.WhereIn("id", PermissionIds).Delete()
		if err != nil {
			return err
		}
		//删除角色按钮
		_, err = rolePermissionModel.WhereIn("permission_id", PermissionIds).Delete()
		if err != nil {
			return err
		}
		//删除角色菜单
		_, err = roleMenuModel.WhereIn("menu_id", req.Id).Delete()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}

func (s sMenu) UpdateSysMenu(ctx context.Context, req *v1.UpdateSysMenuReq) (res *v1.UpdateSysMenuRes, err error) {
	query, err := gjson.EncodeString(req.Query)
	//转换属性
	menu := entity.SysMenu{
		Id:              req.Id,
		ParentId:        req.ParentID,
		MenuType:        req.MenuType,
		MenuName:        req.MenuName,
		I18NKey:         req.I18nKey,
		RouteName:       req.RouteName,
		RoutePath:       req.RoutePath,
		Icon:            req.Icon,
		IconType:        req.IconType,
		Component:       req.Component,
		KeepAlive:       boolToStrPtrIgnoreCase(req.KeepAlive),
		HideInMenu:      boolToStrPtrIgnoreCase(req.HideInMenu),
		Constant:        boolToStrPtrIgnoreCase(req.Constant),
		Href:            req.Href,
		Order:           req.Order,
		MultiTab:        boolToStrPtrIgnoreCase(req.MultiTab),
		FixedIndexInTab: req.FixedIndexInTab,
		ActiveMenu:      req.ActiveMenu,
		Query:           query,
		Status:          req.Status,
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		//要删除的按钮列表
		var sysPermissions []entity.SysPermission
		var deletePermissionIds []int64
		//新的codes(现有的)
		var codes []string
		//菜单按键实体类(现有的)
		buttons := req.Buttons
		permissionModel := dao.SysPermission.Ctx(ctx)
		rolePermissionModel := dao.SysRolePermission.Ctx(ctx)
		menuModel := dao.SysMenu.Ctx(ctx)
		for _, button := range buttons {
			codes = append(codes, button.Code)
		}
		//如果按钮列表为 nil
		if codes == nil {
			req.Buttons = []domain.BTPairs{}
			//删除该菜单下所有的按钮
			err = permissionModel.Where("menu_id", req.Id).Scan(&sysPermissions)
			if err != nil {
				return err
			}
			for _, permission := range sysPermissions {
				deletePermissionIds = append(deletePermissionIds, permission.Id)
			}
			//删除按钮
			_, err = permissionModel.WhereIn("id", deletePermissionIds).Delete()
			if err != nil {
				return err
			}
			//删除按钮关联的数据
			_, err = rolePermissionModel.WhereIn("permission_id", deletePermissionIds).Delete()
			if err != nil {
				return err
			}
		} else {
			//获取需要删除的按钮
			err = permissionModel.Where("menu_id", req.Id).WhereNotIn("code", codes).Scan(&sysPermissions)
			if err != nil {
				return err
			}
			for _, permission := range sysPermissions {
				deletePermissionIds = append(deletePermissionIds, permission.Id)
			}
			//删除按钮
			_, err = permissionModel.WhereIn("id", deletePermissionIds).Delete()
			if err != nil {
				return err
			}
			//删除按钮关联的数据
			_, err = rolePermissionModel.WhereIn("permission_id", deletePermissionIds).Delete()
			if err != nil {
				return err
			}
		}
		//获取需要新增的按钮
		for _, button := range buttons {
			var sysPermission *entity.SysPermission
			err = permissionModel.Unscoped().Where("menu_id", req.Id).Where("code", button.Code).Scan(&sysPermission)
			if err != nil {
				return err
			}
			//如果这个按钮为null
			if sysPermission == nil {
				//添加按钮
				sysPermission = &entity.SysPermission{
					MenuId:      req.Id,
					MenuName:    req.MenuName,
					Code:        button.Code,
					Description: button.Desc,
					Status:      consts.ONE,
				}
				_, err = permissionModel.Insert(sysPermission)
				if err != nil {
					return err
				}
			} else {
				//按键不为nil 则修改按钮状态
				if sysPermission.IsDeleted != consts.ZERO_NUMBER {
					sysPermission.IsDeleted = consts.ZERO_NUMBER
					sysPermission.Description = button.Desc
					_, err = permissionModel.Unscoped().Data(sysPermission).Where("id", sysPermission.Id).Update()
					if err != nil {
						return err
					}
				}
			}
		}
		_, err = menuModel.OmitEmpty().Data(menu).Where("id", req.Id).Update()
		if err != nil {
			return err
		}
		return nil
	})
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

// IsYIgnoreCase 判断布尔是否为True 是则返回Y
func boolToStrPtrIgnoreCase(s bool) string {
	if s {
		return "Y"
	}
	return "N"
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
