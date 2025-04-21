// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysRoleMenuDao is the data access object for the table sys_role_menu.
type SysRoleMenuDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns SysRoleMenuColumns // columns contains all the column names of Table for convenient usage.
}

// SysRoleMenuColumns defines and stores column names for the table sys_role_menu.
type SysRoleMenuColumns struct {
	Id         string //
	RoleId     string // 角色ID
	MenuId     string // 菜单ID
	CreateTime string // 创建时间
	UpdateTime string // 修改时间
	IsDeleted  string // 是否删除(0:否,1:是)
}

// sysRoleMenuColumns holds the columns for the table sys_role_menu.
var sysRoleMenuColumns = SysRoleMenuColumns{
	Id:         "id",
	RoleId:     "role_id",
	MenuId:     "menu_id",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	IsDeleted:  "is_deleted",
}

// NewSysRoleMenuDao creates and returns a new DAO object for table data access.
func NewSysRoleMenuDao() *SysRoleMenuDao {
	return &SysRoleMenuDao{
		group:   "default",
		table:   "sys_role_menu",
		columns: sysRoleMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysRoleMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysRoleMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysRoleMenuDao) Columns() SysRoleMenuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysRoleMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysRoleMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysRoleMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
