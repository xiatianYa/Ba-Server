// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysDictItemDao is the data access object for the table sys_dict_item.
type SysDictItemDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns SysDictItemColumns // columns contains all the column names of Table for convenient usage.
}

// SysDictItemColumns defines and stores column names for the table sys_dict_item.
type SysDictItemColumns struct {
	Id          string // 主键
	DictId      string // 父字典ID
	DictCode    string // 父字典编码
	Value       string // 数据值
	ZhCn        string // 中文名称
	EnUs        string // 英文名称
	Type        string // 类型(前端渲染类型)
	Sort        string // 排序值
	Description string // 字典描述
	CreateTime  string // 创建时间
	UpdateTime  string // 修改时间
	Status      string // 是否启用(0:禁用,1:启用)
	IsDeleted   string // 是否删除(0:否,1:是)
}

// sysDictItemColumns holds the columns for the table sys_dict_item.
var sysDictItemColumns = SysDictItemColumns{
	Id:          "id",
	DictId:      "dict_id",
	DictCode:    "dict_code",
	Value:       "value",
	ZhCn:        "zh_cn",
	EnUs:        "en_us",
	Type:        "type",
	Sort:        "sort",
	Description: "description",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
	Status:      "status",
	IsDeleted:   "is_deleted",
}

// NewSysDictItemDao creates and returns a new DAO object for table data access.
func NewSysDictItemDao() *SysDictItemDao {
	return &SysDictItemDao{
		group:   "default",
		table:   "sys_dict_item",
		columns: sysDictItemColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysDictItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysDictItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysDictItemDao) Columns() SysDictItemColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysDictItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysDictItemDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysDictItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
