package vo

// SysRoleVo is the golang structure for table sys_role.
type SysRoleVo struct {
	Id       int64  `json:"id"`       // 主键
	RoleName string `json:"roleName"` // 角色名称
	RoleCode string `json:"roleCode"` // 角色编码
	RoleDesc string `json:"roleDesc"` // 描述
}
