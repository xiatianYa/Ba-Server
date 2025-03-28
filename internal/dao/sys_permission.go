// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"Ba-Server/internal/dao/internal"
)

// internalSysPermissionDao is an internal type for wrapping the internal DAO implementation.
type internalSysPermissionDao = *internal.SysPermissionDao

// sysPermissionDao is the data access object for the table sys_permission.
// You can define custom methods on it to extend its functionality as needed.
type sysPermissionDao struct {
	internalSysPermissionDao
}

var (
	// SysPermission is a globally accessible object for table sys_permission operations.
	SysPermission = sysPermissionDao{
		internal.NewSysPermissionDao(),
	}
)

// Add your custom methods and functionality below.
