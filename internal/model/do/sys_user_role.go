// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserRole is the golang structure of table sys_user_role for DAO operations like Where/Data.
type SysUserRole struct {
	g.Meta       `orm:"table:sys_user_role, do:true"`
	Id           interface{} //
	UserId       interface{} // 用户ID
	RoleId       interface{} // 角色ID
	CreateUserId interface{} // 创建用户ID
	CreateTime   *gtime.Time // 创建时间
	UpdateUserId interface{} // 修改用户ID
	UpdateTime   *gtime.Time // 修改时间
	Status       interface{} // 是否启用(0:禁用,1:启用)
	IsDeleted    interface{} // 是否删除(0:否,1:是)
}
