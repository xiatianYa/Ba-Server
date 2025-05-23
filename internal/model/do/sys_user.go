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
	g.Meta        `orm:"table:sys_user, do:true"`
	Id            interface{} // 主键ID
	UserName      interface{} // 用户名称
	Password      interface{} // 密码
	NickName      interface{} // 昵称
	UserEmail     interface{} // 邮箱
	UserGender    interface{} // 用户性别
	UserPhone     interface{} // 手机
	CreateTime    *gtime.Time // 创建时间
	UpdateTime    *gtime.Time // 修改时间
	LastLoginTime *gtime.Time // 最后登录时间
	Salt          interface{} // MD5的盐值
	Status        interface{} // 是否启用(0:禁用,1:启用)
	IsDeleted     interface{} // 是否删除(0:否,1:是)
}
