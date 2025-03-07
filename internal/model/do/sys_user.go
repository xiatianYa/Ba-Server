// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure of table sys_user for DAO operations like Where/Data.
type SysUser struct {
	g.Meta             `orm:"table:sys_user, do:true"`
	Id                 interface{} // 主键ID
	UserName           interface{} // 用户名称
	QqOpenId           interface{} // QQ第三方标识
	Password           interface{} // 密码
	NickName           interface{} // 昵称
	Avatar             interface{} // 头像
	UserEmail          interface{} // 邮箱
	UserGender         interface{} // 用户性别
	UserPhone          interface{} // 手机
	CreateUserId       interface{} // 创建用户ID
	CreateTime         *gtime.Time // 创建时间
	UpdateUserId       interface{} // 修改用户ID
	UpdateTime         *gtime.Time // 修改时间
	LastLoginTime      *gtime.Time // 最后登录时间
	UpdatePasswordTime *gtime.Time // 修改密码时间
	Salt               interface{} // MD5的盐值
	Status             interface{} // 是否启用(0:禁用,1:启用)
	IsReset            interface{} // 是否重置
	IsDeleted          interface{} // 是否删除(0:否,1:是)
}
