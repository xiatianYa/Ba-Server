package dto

import "Ba-Server/internal/model/domain"

// SysMenuFormDTO 定义菜单结构体
type SysMenuFormDTO struct {
	// 主键
	Id int64 `json:"id"`
	// 父菜单ID
	ParentID *int64 `json:"parentId"`
	// 菜单类型 1:目录 2:菜单
	MenuType string `json:"menuType"`
	// 菜单名称
	MenuName string `json:"menuName"`
	// 多语言标题
	I18nKey string `json:"i18nKey"`
	// 路由名称
	RouteName string `json:"routeName"`
	// 菜单路径
	RoutePath string `json:"routePath"`
	// 菜单图标
	Icon string `json:"icon"`
	// 图标类型 1:iconify icon 2:local icon
	IconType string `json:"iconType"`
	// 路由组件
	Component string `json:"component"`
	// 缓存页面(Y:是,N:否)
	KeepAlive *bool `json:"keepAlive"`
	// 是否隐藏(Y:是,N:否)
	HideInMenu *bool `json:"hideInMenu"`
	// 是否为常量路由(Y:是,N:否)
	Constant *bool `json:"constant"`
	// 外部链接
	Href string `json:"href"`
	// 排序值
	Order *int `json:"order"`
	// 支持多标签(Y:是,N:否)
	MultiTab *bool `json:"multiTab"`
	// 固定在页签中的序号
	FixedIndexInTab *int `json:"fixedIndexInTab"`
	// 内链URL
	IframeUrl string `json:"iframeUrl"`
	// 是否启用(0:禁用,1:启用)
	Status string `json:"status"`
	// 子菜单
	Children []SysMenuFormDTO `json:"children"`
	// 路由查询参数
	Query []domain.KVPairs `json:"query"`
	// 按钮查询参数
	Buttons []domain.BTPairs `json:"buttons"`
}
