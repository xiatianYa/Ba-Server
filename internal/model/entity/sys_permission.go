// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPermission is the golang structure for table sys_permission.
type SysPermission struct {
	Id          int64       `json:"id"          orm:"id"          ` // 主键
	MenuId      int64       `json:"menuId"      orm:"menu_id"     ` // 菜单ID
	MenuName    string      `json:"menuName"    orm:"menu_name"   ` // 菜单名称
	Code        string      `json:"code"        orm:"code"        ` // 权限资源
	Description string      `json:"description" orm:"description" ` // 描述
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time" ` // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time" ` // 修改时间
	Status      string      `json:"status"      orm:"status"      ` // 是否启用(0:禁用,1:启用)
	IsDeleted   uint        `json:"isDeleted"   orm:"is_deleted"  ` // 是否删除(0:否,1:是)
}
