package route

import (
	"Ba-Server/internal/consts"
	"Ba-Server/internal/dao"
	"Ba-Server/internal/model/domain"
	"Ba-Server/internal/model/entity"
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	sRoute struct{}
)

func init() {
	service.RegisterRoute(New())
}

func New() service.IRoute {
	return &sRoute{}
}

func (s sRoute) GetConstantRoutes(ctx context.Context) (constantRoutes []map[string]interface{}, err error) {
	data := []map[string]interface{}{
		{
			"name":      "403",
			"path":      "/403",
			"component": "layout.blank$view.403",
			"props":     false,
			"meta": map[string]interface{}{
				"title":           "403",
				"i18nKey":         "route.403",
				"localIcon":       "",
				"icon":            "",
				"order":           0,
				"hideInMenu":      true,
				"activeMenu":      "",
				"multiTab":        false,
				"fixedIndexInTab": nil,
				"keepAlive":       false,
				"constant":        true,
				"href":            "",
				"query":           nil,
			},
		},
		{
			"name":      "404",
			"path":      "/404",
			"component": "layout.blank$view.404",
			"props":     false,
			"meta": map[string]interface{}{
				"title":           "404",
				"i18nKey":         "route.404",
				"localIcon":       "",
				"icon":            "",
				"order":           0,
				"hideInMenu":      true,
				"activeMenu":      "",
				"multiTab":        false,
				"fixedIndexInTab": nil,
				"keepAlive":       false,
				"constant":        true,
				"href":            "",
				"query":           nil,
			},
		},
		{
			"name":      "500",
			"path":      "/500",
			"component": "layout.blank$view.500",
			"props":     false,
			"meta": map[string]interface{}{
				"title":           "500",
				"i18nKey":         "route.500",
				"localIcon":       "",
				"icon":            "",
				"order":           0,
				"hideInMenu":      true,
				"activeMenu":      "",
				"multiTab":        false,
				"fixedIndexInTab": nil,
				"keepAlive":       false,
				"constant":        true,
				"href":            "",
				"query":           nil,
			},
		},
		{
			"name":      "login",
			"path":      "/login/:module(pwd-login|code-login|register|reset-pwd|bind-wechat)?",
			"component": "layout.blank$view.login",
			"props":     true,
			"meta": map[string]interface{}{
				"title":           "login",
				"i18nKey":         "route.login",
				"localIcon":       "",
				"icon":            "",
				"order":           0,
				"hideInMenu":      true,
				"activeMenu":      "",
				"multiTab":        false,
				"fixedIndexInTab": nil,
				"keepAlive":       false,
				"constant":        true,
				"href":            "",
				"query":           nil,
			},
		},
	}
	return data, nil
}

func (s sRoute) GetUserRoutes(ctx context.Context) (sysUserInfoVo *vo.SysUserRouteVO, err error) {
	//获取用户ID
	userId := ctx.Value(consts.UserId)
	if userId == nil {
		return nil, gerror.New("用户ID不存在")
	}
	// 返回的路由结构
	var resultRoute []domain.Route
	// 用户用户的角色对应列表
	var sysUserRoles []entity.SysUserRole
	// 用户拥有的角色列表
	var sysRoles []entity.SysRole
	// 用户用户菜单列表
	var sysRoleMenus []entity.SysRoleMenu
	// 创建一个新的切片来存储 roleIds
	var roleIds []int64
	// 创建一个新的切片来存储 menuIds
	var menuIds []int64
	// parentMenus 父菜单
	var parentMenus []entity.SysMenu
	//创建用户权限模型
	roleModel := dao.SysRole.Ctx(ctx)
	userRoleModel := dao.SysUserRole.Ctx(ctx)
	roleMenuModel := dao.SysRoleMenu.Ctx(ctx)
	menuModel := dao.SysMenu.Ctx(ctx)
	//查询用户角色
	_ = userRoleModel.Where("user_id=?", 1).Scan(&sysUserRoles)
	// 遍历 sysUserRoles 切片，将每个 Role 结构体的 roleId 字段添加到 roleIds 切片中
	for _, role := range sysUserRoles {
		roleIds = append(roleIds, role.RoleId)
	}
	_ = roleModel.WhereIn("id", roleIds).Scan(&sysRoles)
	//查询用户菜单
	_ = roleMenuModel.WhereIn("role_id", roleIds).Scan(&sysRoleMenus)
	//遍历 sysRoleMenu 切片，将每个 Menu 结构体的 menuId 字段添加到 menuIds 切片中
	for _, roleMenu := range sysRoleMenus {
		menuIds = append(menuIds, roleMenu.MenuId)
	}
	//查询所有符合条件的父菜单
	_ = menuModel.Where("status=?", consts.ONE_NUMBER).Where("parent_id=?", consts.ZERO_NUMBER).WhereIn("id", menuIds).Scan(&parentMenus)
	//遍历有权限的父菜单 递归构建菜单树
	for _, parentMenu := range parentMenus {
		route := buildMenuTree(menuModel, parentMenu, menuIds)
		resultRoute = append(resultRoute, route)
	}
	return &vo.SysUserRouteVO{
		Home:   "home",
		Routes: resultRoute,
	}, nil
}

// buildMenuTree 递归构建菜单
func buildMenuTree(menuModel *gdb.Model, sysMenu entity.SysMenu, menuIds []int64) domain.Route {
	// children 子菜单
	var childrenMenu []entity.SysMenu
	route := SysUserRouteVOBuilder(sysMenu, childrenMenu)
	//查询该菜单下的子菜单
	_ = menuModel.Where("status=?", consts.ONE_NUMBER).Where("parent_id=?", sysMenu.Id).WhereIn("id", menuIds).Scan(&childrenMenu)
	//构建子菜单到父菜单里
	for _, menu := range childrenMenu {
		children := buildMenuTree(menuModel, menu, menuIds)
		route.Children = append(route.Children, children)
	}
	return route
}

// buildTree 构建树结构
func buildTree(menus []entity.SysMenu, rootID int64) []domain.Route {
	tree := make([]domain.Route, 0)
	nodeMap := make(map[int64]domain.Route)

	// 初始化节点映射
	for _, menu := range menus {
		var query []domain.KVPairs
		if menu.Query != "" {
			_ = json.Unmarshal([]byte(menu.Query), &query)
		}
		route := domain.Route{
			Name:      menu.RouteName,
			Path:      menu.RoutePath,
			Component: menu.Component,
			Meta: domain.Meta{
				Title:           menu.RouteName,
				I18nKey:         menu.I18NKey,
				Order:           menu.Order,
				Href:            menu.Href,
				FixedIndexInTab: getFixedIndexInTab(menu.FixedIndexInTab),
				Query:           query,
				ActiveMenu:      menu.ActiveMenu,
				KeepAlive:       menu.KeepAlive == "Y",
				Constant:        menu.Constant == "Y",
				HideInMenu:      menu.HideInMenu == "Y",
				MultiTab:        menu.MultiTab == "Y",
				Icon:            getIcon(menu.IconType, menu.Icon),
				LocalIcon:       getLocalIcon(menu.IconType, menu.Icon),
			},
			Children: make([]domain.Route, 0),
		}
		nodeMap[menu.Id] = route
	}

	// 构建树结构
	for _, menu := range menus {
		if menu.ParentId == rootID {
			tree = append(tree, nodeMap[menu.Id])
		} else {
			if parent, ok := nodeMap[menu.ParentId]; ok {
				parent.Children = append(parent.Children, nodeMap[menu.Id])
				nodeMap[menu.ParentId] = parent
			}
		}
	}

	return tree
}

func SysUserRouteVOBuilder(sysMenu entity.SysMenu, sysMenus []entity.SysMenu) domain.Route {
	sysMenus = append(sysMenus, sysMenu)

	// 构建树结构
	tree := buildTree(sysMenus, sysMenu.Id)
	var query []domain.KVPairs
	if sysMenu.Query != "" {
		_ = json.Unmarshal([]byte(sysMenu.Query), &query)
	}

	return domain.Route{
		Name:      sysMenu.RouteName,
		Path:      sysMenu.RoutePath,
		Component: sysMenu.Component,
		Meta: domain.Meta{
			Title:           sysMenu.RouteName,
			I18nKey:         sysMenu.I18NKey,
			Order:           sysMenu.Order,
			Href:            sysMenu.Href,
			FixedIndexInTab: getFixedIndexInTab(sysMenu.FixedIndexInTab),
			Query:           query,
			ActiveMenu:      sysMenu.ActiveMenu,
			KeepAlive:       sysMenu.KeepAlive == "Y",
			Constant:        sysMenu.Constant == "Y",
			HideInMenu:      sysMenu.HideInMenu == "Y",
			MultiTab:        sysMenu.MultiTab == "Y",
			Icon:            getIcon(sysMenu.IconType, sysMenu.Icon),
			LocalIcon:       getLocalIcon(sysMenu.IconType, sysMenu.Icon),
		},
		Children: tree,
	}
}

// getIcon 获取图标
func getIcon(iconType, icon string) string {
	if iconType == "1" {
		return icon
	}
	return ""
}

// getIcon 获取固定在页签中的序号
func getFixedIndexInTab(fixedIndexInTab int) *int {
	if fixedIndexInTab == 1 {
		return &fixedIndexInTab
	}
	return nil
}

// FixedIndexInTab 获取本地图标
func getLocalIcon(iconType, icon string) string {
	if iconType == "2" {
		return icon
	}
	return ""
}
