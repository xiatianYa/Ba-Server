// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"Ba-Server/internal/dao/internal"
)

// internalSysRoleMenuDao is an internal type for wrapping the internal DAO implementation.
type internalSysRoleMenuDao = *internal.SysRoleMenuDao

// sysRoleMenuDao is the data access object for the table sys_role_menu.
// You can define custom methods on it to extend its functionality as needed.
type sysRoleMenuDao struct {
	internalSysRoleMenuDao
}

var (
	// SysRoleMenu is a globally accessible object for table sys_role_menu operations.
	SysRoleMenu = sysRoleMenuDao{
		internal.NewSysRoleMenuDao(),
	}
)

// Add your custom methods and functionality below.
