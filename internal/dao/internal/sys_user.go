// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserDao is the data access object for the table sys_user.
type SysUserDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of the current DAO.
	columns SysUserColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserColumns defines and stores column names for the table sys_user.
type SysUserColumns struct {
	Id            string // 主键ID
	UserName      string // 用户名称
	Password      string // 密码
	NickName      string // 昵称
	UserEmail     string // 邮箱
	UserGender    string // 用户性别
	UserPhone     string // 手机
	CreateTime    string // 创建时间
	UpdateTime    string // 修改时间
	LastLoginTime string // 最后登录时间
	Salt          string // MD5的盐值
	Status        string // 是否启用(0:禁用,1:启用)
	IsDeleted     string // 是否删除(0:否,1:是)
}

// sysUserColumns holds the columns for the table sys_user.
var sysUserColumns = SysUserColumns{
	Id:            "id",
	UserName:      "user_name",
	Password:      "password",
	NickName:      "nick_name",
	UserEmail:     "user_email",
	UserGender:    "user_gender",
	UserPhone:     "user_phone",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
	LastLoginTime: "last_login_time",
	Salt:          "salt",
	Status:        "status",
	IsDeleted:     "is_deleted",
}

// NewSysUserDao creates and returns a new DAO object for table data access.
func NewSysUserDao() *SysUserDao {
	return &SysUserDao{
		group:   "default",
		table:   "sys_user",
		columns: sysUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysUserDao) Columns() SysUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
