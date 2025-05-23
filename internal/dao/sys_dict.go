// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"Ba-Server/internal/dao/internal"
)

// internalSysDictDao is an internal type for wrapping the internal DAO implementation.
type internalSysDictDao = *internal.SysDictDao

// sysDictDao is the data access object for the table sys_dict.
// You can define custom methods on it to extend its functionality as needed.
type sysDictDao struct {
	internalSysDictDao
}

var (
	// SysDict is a globally accessible object for table sys_dict operations.
	SysDict = sysDictDao{
		internal.NewSysDictDao(),
	}
)

// Add your custom methods and functionality below.
