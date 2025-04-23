// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MonLogsOperationDao is the data access object for the table mon_logs_operation.
type MonLogsOperationDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns MonLogsOperationColumns // columns contains all the column names of Table for convenient usage.
}

// MonLogsOperationColumns defines and stores column names for the table mon_logs_operation.
type MonLogsOperationColumns struct {
	Id            string // 主键ID
	Ip            string // IP
	IpAddr        string // IP所属地
	UserAgent     string // 登录代理
	RequestUri    string // 请求URI
	RequestPath   string // 请求PATH
	RequestMethod string // 请求方式
	MethodParams  string // 请求参数
	UseTime       string // 请求耗时
	CreateTime    string // 创建时间
	UpdateTime    string // 修改时间
	IsDeleted     string // 是否删除(0:否,1:是)
}

// monLogsOperationColumns holds the columns for the table mon_logs_operation.
var monLogsOperationColumns = MonLogsOperationColumns{
	Id:            "id",
	Ip:            "ip",
	IpAddr:        "ip_addr",
	UserAgent:     "user_agent",
	RequestUri:    "request_uri",
	RequestPath:   "request_path",
	RequestMethod: "request_method",
	MethodParams:  "method_params",
	UseTime:       "use_time",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
	IsDeleted:     "is_deleted",
}

// NewMonLogsOperationDao creates and returns a new DAO object for table data access.
func NewMonLogsOperationDao() *MonLogsOperationDao {
	return &MonLogsOperationDao{
		group:   "default",
		table:   "mon_logs_operation",
		columns: monLogsOperationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MonLogsOperationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MonLogsOperationDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MonLogsOperationDao) Columns() MonLogsOperationColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MonLogsOperationDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MonLogsOperationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *MonLogsOperationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
