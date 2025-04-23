// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserRole is the golang structure for table sys_user_role.
type SysUserRole struct {
	Id         int64       `json:"id"         orm:"id"          ` //
	UserId     int64       `json:"userId"     orm:"user_id"     ` // 用户ID
	RoleId     int64       `json:"roleId"     orm:"role_id"     ` // 角色ID
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" ` // 修改时间
	IsDeleted  uint        `json:"isDeleted"  orm:"is_deleted"  ` // 是否删除(0:否,1:是)
}
