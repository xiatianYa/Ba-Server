// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure of table sys_menu for DAO operations like Where/Data.
type SysMenu struct {
	g.Meta          `orm:"table:sys_menu, do:true"`
	Id              interface{} // 主键
	ParentId        interface{} // 父菜单ID
	MenuType        interface{} // 菜单类型 1:目录 2:菜单
	MenuName        interface{} // 菜单名称
	I18NKey         interface{} // 多语言标题
	RouteName       interface{} // 路由名称
	RoutePath       interface{} // 菜单路径
	Icon            interface{} // 菜单图标
	IconType        interface{} // 图标类型 1:iconify icon 2:local icon
	Component       interface{} // 路由组件
	KeepAlive       interface{} // 缓存页面(Y:是,N:否)
	HideInMenu      interface{} // 是否隐藏(Y:是,N:否)
	Constant        interface{} // 是否为常量路由(Y:是,N:否)
	Href            interface{} // 外部链接
	Order           interface{} // 排序值
	MultiTab        interface{} // 支持多标签(Y:是,N:否)
	FixedIndexInTab interface{} // 固定在页签中的序号
	ActiveMenu      interface{} // 内链URL
	Query           interface{} // 路由查询参数
	CreateTime      *gtime.Time // 创建时间
	UpdateTime      *gtime.Time // 修改时间
	Status          interface{} // 是否启用(0:禁用,1:启用)
	IsDeleted       interface{} // 是否删除(0:否,1:是)
}
