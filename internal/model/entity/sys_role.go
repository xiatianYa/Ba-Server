// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Id         int64       `json:"id"         orm:"id"          ` // 主键
	RoleName   string      `json:"roleName"   orm:"role_name"   ` // 角色名称
	RoleCode   string      `json:"roleCode"   orm:"role_code"   ` // 角色编码
	RoleDesc   string      `json:"roleDesc"   orm:"role_desc"   ` // 描述
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" ` // 修改时间
	Status     string      `json:"status"     orm:"status"      ` // 是否启用(0:禁用,1:启用)
	IsDeleted  uint        `json:"isDeleted"  orm:"is_deleted"  ` // 是否删除(0:否,1:是)
}
