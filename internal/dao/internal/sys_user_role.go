// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserRoleDao is the data access object for the table sys_user_role.
type SysUserRoleDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns SysUserRoleColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserRoleColumns defines and stores column names for the table sys_user_role.
type SysUserRoleColumns struct {
	Id           string //
	UserId       string // 用户ID
	RoleId       string // 角色ID
	CreateUserId string // 创建用户ID
	CreateTime   string // 创建时间
	UpdateUserId string // 修改用户ID
	UpdateTime   string // 修改时间
	Status       string // 是否启用(0:禁用,1:启用)
	IsDeleted    string // 是否删除(0:否,1:是)
}

// sysUserRoleColumns holds the columns for the table sys_user_role.
var sysUserRoleColumns = SysUserRoleColumns{
	Id:           "id",
	UserId:       "user_id",
	RoleId:       "role_id",
	CreateUserId: "create_user_id",
	CreateTime:   "create_time",
	UpdateUserId: "update_user_id",
	UpdateTime:   "update_time",
	Status:       "status",
	IsDeleted:    "is_deleted",
}

// NewSysUserRoleDao creates and returns a new DAO object for table data access.
func NewSysUserRoleDao() *SysUserRoleDao {
	return &SysUserRoleDao{
		group:   "default",
		table:   "sys_user_role",
		columns: sysUserRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysUserRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysUserRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysUserRoleDao) Columns() SysUserRoleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysUserRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysUserRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysUserRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
