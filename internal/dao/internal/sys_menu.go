// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysMenuDao is the data access object for the table sys_menu.
type SysMenuDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns SysMenuColumns // columns contains all the column names of Table for convenient usage.
}

// SysMenuColumns defines and stores column names for the table sys_menu.
type SysMenuColumns struct {
	Id              string // 主键
	ParentId        string // 父菜单ID
	MenuType        string // 菜单类型 1:目录 2:菜单
	MenuName        string // 菜单名称
	I18NKey         string // 多语言标题
	RouteName       string // 路由名称
	RoutePath       string // 菜单路径
	Icon            string // 菜单图标
	IconType        string // 图标类型 1:iconify icon 2:local icon
	Component       string // 路由组件
	KeepAlive       string // 缓存页面(Y:是,N:否)
	HideInMenu      string // 是否隐藏(Y:是,N:否)
	Constant        string // 是否为常量路由(Y:是,N:否)
	Href            string // 外部链接
	Order           string // 排序值
	MultiTab        string // 支持多标签(Y:是,N:否)
	FixedIndexInTab string // 固定在页签中的序号
	ActiveMenu      string // 内链URL
	Query           string // 路由查询参数
	CreateUserId    string // 创建用户ID
	CreateTime      string // 创建时间
	UpdateUserId    string // 修改用户ID
	UpdateTime      string // 修改时间
	Status          string // 是否启用(0:禁用,1:启用)
	IsDeleted       string // 是否删除(0:否,1:是)
}

// sysMenuColumns holds the columns for the table sys_menu.
var sysMenuColumns = SysMenuColumns{
	Id:              "id",
	ParentId:        "parent_id",
	MenuType:        "menu_type",
	MenuName:        "menu_name",
	I18NKey:         "i18n_key",
	RouteName:       "route_name",
	RoutePath:       "route_path",
	Icon:            "icon",
	IconType:        "icon_type",
	Component:       "component",
	KeepAlive:       "keep_alive",
	HideInMenu:      "hide_in_menu",
	Constant:        "constant",
	Href:            "href",
	Order:           "order",
	MultiTab:        "multi_tab",
	FixedIndexInTab: "fixed_index_in_tab",
	ActiveMenu:      "active_menu",
	Query:           "query",
	CreateUserId:    "create_user_id",
	CreateTime:      "create_time",
	UpdateUserId:    "update_user_id",
	UpdateTime:      "update_time",
	Status:          "status",
	IsDeleted:       "is_deleted",
}

// NewSysMenuDao creates and returns a new DAO object for table data access.
func NewSysMenuDao() *SysMenuDao {
	return &SysMenuDao{
		group:   "default",
		table:   "sys_menu",
		columns: sysMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysMenuDao) Columns() SysMenuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
