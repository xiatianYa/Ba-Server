// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPermission is the golang structure of table sys_permission for DAO operations like Where/Data.
type SysPermission struct {
	g.Meta      `orm:"table:sys_permission, do:true"`
	Id          interface{} // 主键
	MenuId      interface{} // 菜单ID
	MenuName    interface{} // 菜单名称
	Code        interface{} // 权限资源
	Description interface{} // 描述
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 修改时间
	Status      interface{} // 是否启用(0:禁用,1:启用)
	IsDeleted   interface{} // 是否删除(0:否,1:是)
}
