// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDictDao is the data access object for the table sys_dict.
type SysDictDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns SysDictColumns // columns contains all the column names of Table for convenient usage.
}

// SysDictColumns defines and stores column names for the table sys_dict.
type SysDictColumns struct {
	Id           string // 主键
	Name         string // 字典名称
	Code         string // 字典编码
	Type         string // 字典类型(1:系统字典,2:业务字典)
	Sort         string // 排序值
	Description  string // 字典描述
	CreateUserId string // 创建用户ID
	CreateTime   string // 创建时间
	UpdateUserId string // 修改用户ID
	UpdateTime   string // 修改时间
	Status       string // 是否启用(0:禁用,1:启用)
	IsDeleted    string // 是否删除(0:否,1:是)
}

// sysDictColumns holds the columns for the table sys_dict.
var sysDictColumns = SysDictColumns{
	Id:           "id",
	Name:         "name",
	Code:         "code",
	Type:         "type",
	Sort:         "sort",
	Description:  "description",
	CreateUserId: "create_user_id",
	CreateTime:   "create_time",
	UpdateUserId: "update_user_id",
	UpdateTime:   "update_time",
	Status:       "status",
	IsDeleted:    "is_deleted",
}

// NewSysDictDao creates and returns a new DAO object for table data access.
func NewSysDictDao() *SysDictDao {
	return &SysDictDao{
		group:   "default",
		table:   "sys_dict",
		columns: sysDictColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysDictDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysDictDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysDictDao) Columns() SysDictColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysDictDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysDictDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysDictDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
