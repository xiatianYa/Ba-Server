// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure of table sys_role for DAO operations like Where/Data.
type SysRole struct {
	g.Meta       `orm:"table:sys_role, do:true"`
	Id           interface{} // 主键
	RoleName     interface{} // 角色名称
	RoleCode     interface{} // 角色编码
	RoleDesc     interface{} // 描述
	CreateUserId interface{} // 创建用户ID
	CreateTime   *gtime.Time // 创建时间
	UpdateUserId interface{} // 修改用户ID
	UpdateTime   *gtime.Time // 修改时间
	Status       interface{} // 是否启用(0:禁用,1:启用)
	IsDeleted    interface{} // 是否删除(0:否,1:是)
}
