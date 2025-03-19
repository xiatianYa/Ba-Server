package v1

import (
	"Ba-Server/internal/model/domain"
	"Ba-Server/internal/model/vo"
	"github.com/gogf/gf/v2/frame/g"
)

type GetAllSysMenusReq struct {
	g.Meta `path:"/getAllMenus" tags:"getAllMenus" method:"get" summary:"获取菜单名称列表"`
}

type GetAllSysMenusRes []string

type GetSysMenuTreeReq struct {
	g.Meta `path:"/getMenuTree" tags:"getMenuTree" method:"get" summary:"获取菜单树"`
}

type GetSysMenuTreeRes []vo.SysMenuTreeVo

type GetSysMenuIdsByRoleIdReq struct {
	g.Meta `path:"/getMenuIdsByRoleId/{id}" tags:"getMenuIdsByRoleId" method:"get" summary:"获取菜单Ids根据角色Id"`
	Id     int64 `v:"required" dc:"role id"`
}

type GetSysMenuIdsByRoleIdRes []int64

type GetSysMenuPageReq struct {
	g.Meta  `path:"/page" tags:"page" method:"get" summary:"获取菜单列表分页"`
	Current int `v:"required" json:"current"`
	Size    int `v:"required" json:"size"`
}

type GetSysMenuPageRes *domain.RPage

// SaveSysMenuReq 定义添加菜单的请求结构体
type SaveSysMenuReq struct {
	g.Meta `path:"/save" tags:"save" method:"post" summary:"添加菜单"`
	// 父菜单ID
	ParentID int64 `json:"parentId" dc:"父菜单ID"`
	// 菜单类型 1:目录 2:菜单
	MenuType string `json:"menuType" dc:"菜单类型 1:目录 2:菜单"`
	// 菜单名称
	MenuName string `json:"menuName" dc:"菜单名称"`
	// 多语言标题
	I18nKey string `json:"i18nKey" dc:"多语言标题"`
	// 路由名称
	RouteName string `json:"routeName" dc:"路由名称"`
	// 菜单路径
	RoutePath string `json:"routePath" dc:"菜单路径"`
	// 菜单图标
	Icon string `json:"icon" dc:"菜单图标"`
	// 图标类型 1:iconify icon 2:local icon
	IconType string `json:"iconType" dc:"图标类型 1:iconify icon 2:local icon"`
	// 路由组件
	Component string `json:"component" dc:"路由组件"`
	// 缓存页面(Y:是,N:否)
	KeepAlive bool `json:"keepAlive" dc:"缓存页面(Y:是,N:否)"`
	// 是否隐藏(Y:是,N:否)
	HideInMenu bool `json:"hideInMenu" dc:"是否隐藏(Y:是,N:否)"`
	// 是否为常量路由(Y:是,N:否)
	Constant bool `json:"constant" dc:"是否为常量路由(Y:是,N:否)"`
	// 外部链接
	Href string `json:"href" dc:"外部链接"`
	// 排序值
	Order int `json:"order" dc:"排序值"`
	// 支持多标签(Y:是,N:否)
	MultiTab bool `json:"multiTab" dc:"支持多标签(Y:是,N:否)"`
	// 固定在页签中的序号
	FixedIndexInTab int `json:"fixedIndexInTab" dc:"固定在页签中的序号"`
	// 内链URL
	IframeUrl string `json:"iframeUrl" dc:"内链URL"`
	// 是否启用(0:禁用,1:启用)
	Status string `json:"status" dc:"是否启用(0:禁用,1:启用)"`
	// 路由查询参数
	Query []domain.KVPairs `json:"query" dc:"路由查询参数"`
	// 按钮查询参数
	Buttons []domain.BTPairs `json:"buttons" dc:"按钮查询参数"`
}

type SaveSysMenuRes bool
