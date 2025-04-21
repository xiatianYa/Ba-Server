// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure for table sys_menu.
type SysMenu struct {
	Id              int64       `json:"id"              orm:"id"                 ` // 主键
	ParentId        int64       `json:"parentId"        orm:"parent_id"          ` // 父菜单ID
	MenuType        string      `json:"menuType"        orm:"menu_type"          ` // 菜单类型 1:目录 2:菜单
	MenuName        string      `json:"menuName"        orm:"menu_name"          ` // 菜单名称
	I18NKey         string      `json:"i18NKey"         orm:"i18n_key"           ` // 多语言标题
	RouteName       string      `json:"routeName"       orm:"route_name"         ` // 路由名称
	RoutePath       string      `json:"routePath"       orm:"route_path"         ` // 菜单路径
	Icon            string      `json:"icon"            orm:"icon"               ` // 菜单图标
	IconType        string      `json:"iconType"        orm:"icon_type"          ` // 图标类型 1:iconify icon 2:local icon
	Component       string      `json:"component"       orm:"component"          ` // 路由组件
	KeepAlive       string      `json:"keepAlive"       orm:"keep_alive"         ` // 缓存页面(Y:是,N:否)
	HideInMenu      string      `json:"hideInMenu"      orm:"hide_in_menu"       ` // 是否隐藏(Y:是,N:否)
	Constant        string      `json:"constant"        orm:"constant"           ` // 是否为常量路由(Y:是,N:否)
	Href            string      `json:"href"            orm:"href"               ` // 外部链接
	Order           int         `json:"order"           orm:"order"              ` // 排序值
	MultiTab        string      `json:"multiTab"        orm:"multi_tab"          ` // 支持多标签(Y:是,N:否)
	FixedIndexInTab int         `json:"fixedIndexInTab" orm:"fixed_index_in_tab" ` // 固定在页签中的序号
	ActiveMenu      string      `json:"activeMenu"      orm:"active_menu"        ` // 内链URL
	Query           string      `json:"query"           orm:"query"              ` // 路由查询参数
	CreateTime      *gtime.Time `json:"createTime"      orm:"create_time"        ` // 创建时间
	UpdateTime      *gtime.Time `json:"updateTime"      orm:"update_time"        ` // 修改时间
	Status          string      `json:"status"          orm:"status"             ` // 是否启用(0:禁用,1:启用)
	IsDeleted       uint        `json:"isDeleted"       orm:"is_deleted"         ` // 是否删除(0:否,1:是)
}
