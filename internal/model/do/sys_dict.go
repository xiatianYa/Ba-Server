// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDict is the golang structure of table sys_dict for DAO operations like Where/Data.
type SysDict struct {
	g.Meta       `orm:"table:sys_dict, do:true"`
	Id           interface{} // 主键
	Name         interface{} // 字典名称
	Code         interface{} // 字典编码
	Type         interface{} // 字典类型(1:系统字典,2:业务字典)
	Sort         interface{} // 排序值
	Description  interface{} // 字典描述
	CreateUserId interface{} // 创建用户ID
	CreateTime   *gtime.Time // 创建时间
	UpdateUserId interface{} // 修改用户ID
	UpdateTime   *gtime.Time // 修改时间
	Status       interface{} // 是否启用(0:禁用,1:启用)
	IsDeleted    interface{} // 是否删除(0:否,1:是)
}
