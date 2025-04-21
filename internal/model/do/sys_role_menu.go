// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRoleMenu is the golang structure of table sys_role_menu for DAO operations like Where/Data.
type SysRoleMenu struct {
	g.Meta     `orm:"table:sys_role_menu, do:true"`
	Id         interface{} //
	RoleId     interface{} // 角色ID
	MenuId     interface{} // 菜单ID
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 修改时间
	IsDeleted  interface{} // 是否删除(0:否,1:是)
}
