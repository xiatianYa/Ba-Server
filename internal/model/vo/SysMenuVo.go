package vo

import "Ba-Server/internal/model/domain"

// SysMenuVo is the golang structure for table sys_menu.
type SysMenuVo struct {
	// 主键
	Id int64 `json:"id"`
	// 父菜单 ID
	ParentId int64 `json:"parentId"`
	// 菜单类型 1:目录 2:菜单
	MenuType string `json:"menuType"`
	// 菜单名称
	MenuName string `json:"menuName"`
	// 多语言标题
	I18nKey string `json:"i18nKey"`
	// 路由名称
	RouteName string `json:"routeName"`
	// 路由路径
	RoutePath string `json:"routePath"`
	// 菜单图标
	Icon string `json:"icon"`
	// 图标类型 1:iconic icon 2:local icon
	IconType string `json:"iconType"`
	// 路由组件
	Component string `json:"component"`
	// 缓存路由(Y:是,N:否)
	KeepAlive *bool `json:"keepAlive"`
	// 是否隐藏(Y:是,N:否)
	HideInMenu *bool `json:"hideInMenu"`
	// 是否为常量路由(Y:是,N:否)
	Constant *bool `json:"constant"`
	// 外部链接
	Href string `json:"href"`
	// 内嵌链接 Iframe URL
	ActiveMenu string `json:"activeMenu"`
	// 排序值
	Order int `json:"order"`
	// 支持多标签(Y:是,N:否)
	MultiTab *bool `json:"multiTab"`
	// 固定在页签中的序号
	FixedIndexInTab int `json:"fixedIndexInTab"`
	// 路由查询参数 JSON 字符串
	Query []domain.KVPairs `json:"query"`
	// 路由查询参数 JSON 字符串
	Buttons []domain.BTPairs `json:"buttons"`
	// 是否启用(0:禁用,1:启用)
	Status string `json:"status"`
	// 子对象
	Children []SysMenuVo `json:"children"`
}
