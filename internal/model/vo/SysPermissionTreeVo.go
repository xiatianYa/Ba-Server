package vo

// SysPermissionTreeVo is the golang structure for table sys_menu.
type SysPermissionTreeVo struct {
	Label    string                `json:"label"`    // 权限名称
	Value    int64                 `json:"value"`    // 权限值
	Children []SysPermissionTreeVo `json:"children"` //字权限树
}
