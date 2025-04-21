// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRolePermission is the golang structure for table sys_role_permission.
type SysRolePermission struct {
	Id           int64       `json:"id"           orm:"id"            ` //
	RoleId       int64       `json:"roleId"       orm:"role_id"       ` // 角色ID
	PermissionId int64       `json:"permissionId" orm:"permission_id" ` // 权限ID
	CreateTime   *gtime.Time `json:"createTime"   orm:"create_time"   ` // 创建时间
	UpdateTime   *gtime.Time `json:"updateTime"   orm:"update_time"   ` // 修改时间
	IsDeleted    uint        `json:"isDeleted"    orm:"is_deleted"    ` // 是否删除(0:否,1:是)
}
