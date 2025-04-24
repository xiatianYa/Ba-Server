// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MonLogsFileDao is the data access object for the table mon_logs_file.
type MonLogsFileDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns MonLogsFileColumns // columns contains all the column names of Table for convenient usage.
}

// MonLogsFileColumns defines and stores column names for the table mon_logs_file.
type MonLogsFileColumns struct {
	Id           string // 主键ID
	UserId       string // 用户ID
	FilePath     string // 文件物理路径
	FileUrl      string // 文件网络路径
	FileType     string // 文件类型
	FileSize     string // 文件大小
	ErrorMsg     string // 异常日志
	Status       string // 上传状态(0:失败 1:成功)
	CreateUserId string // 创建用户ID
	CreateTime   string // 创建时间
	UpdateUserId string // 修改用户ID
	UpdateTime   string // 修改时间
	IsDeleted    string // 是否删除(0:否,1:是)
}

// monLogsFileColumns holds the columns for the table mon_logs_file.
var monLogsFileColumns = MonLogsFileColumns{
	Id:           "id",
	UserId:       "user_id",
	FilePath:     "file_path",
	FileUrl:      "file_url",
	FileType:     "file_type",
	FileSize:     "file_size",
	ErrorMsg:     "error_msg",
	Status:       "status",
	CreateUserId: "create_user_id",
	CreateTime:   "create_time",
	UpdateUserId: "update_user_id",
	UpdateTime:   "update_time",
	IsDeleted:    "is_deleted",
}

// NewMonLogsFileDao creates and returns a new DAO object for table data access.
func NewMonLogsFileDao() *MonLogsFileDao {
	return &MonLogsFileDao{
		group:   "default",
		table:   "mon_logs_file",
		columns: monLogsFileColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MonLogsFileDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MonLogsFileDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MonLogsFileDao) Columns() MonLogsFileColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MonLogsFileDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MonLogsFileDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *MonLogsFileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
