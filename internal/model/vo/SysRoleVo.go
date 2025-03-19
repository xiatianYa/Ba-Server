package vo

import "github.com/gogf/gf/v2/os/gtime"

// SysRoleVo is the golang structure for table sys_role.
type SysRoleVo struct {
	Id         int64       `json:"id"`         // 主键
	RoleName   string      `json:"roleName"`   // 角色名称
	RoleCode   string      `json:"roleCode"`   // 角色编码
	RoleDesc   string      `json:"roleDesc"`   // 描述
	CreateTime *gtime.Time `json:"createTime"` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime"` // 修改时间
	Status     string      `json:"status"`     // 是否启用(0:禁用,1:启用)
	IsDeleted  uint        `json:"isDeleted"`  // 是否删除(0:否,1:是)
}
