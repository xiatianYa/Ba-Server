// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysPermissionDao is the data access object for the table sys_permission.
type SysPermissionDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns SysPermissionColumns // columns contains all the column names of Table for convenient usage.
}

// SysPermissionColumns defines and stores column names for the table sys_permission.
type SysPermissionColumns struct {
	Id          string // 主键
	MenuId      string // 菜单ID
	MenuName    string // 菜单名称
	Code        string // 权限资源
	Description string // 描述
	CreateTime  string // 创建时间
	UpdateTime  string // 修改时间
	Status      string // 是否启用(0:禁用,1:启用)
	IsDeleted   string // 是否删除(0:否,1:是)
}

// sysPermissionColumns holds the columns for the table sys_permission.
var sysPermissionColumns = SysPermissionColumns{
	Id:          "id",
	MenuId:      "menu_id",
	MenuName:    "menu_name",
	Code:        "code",
	Description: "description",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
	Status:      "status",
	IsDeleted:   "is_deleted",
}

// NewSysPermissionDao creates and returns a new DAO object for table data access.
func NewSysPermissionDao() *SysPermissionDao {
	return &SysPermissionDao{
		group:   "default",
		table:   "sys_permission",
		columns: sysPermissionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysPermissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysPermissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysPermissionDao) Columns() SysPermissionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysPermissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysPermissionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysPermissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
