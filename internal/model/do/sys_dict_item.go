// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictItem is the golang structure of table sys_dict_item for DAO operations like Where/Data.
type SysDictItem struct {
	g.Meta       `orm:"table:sys_dict_item, do:true"`
	Id           interface{} // 主键
	DictId       interface{} // 父字典ID
	DictCode     interface{} // 父字典编码
	Value        interface{} // 数据值
	ZhCn         interface{} // 中文名称
	EnUs         interface{} // 英文名称
	Type         interface{} // 类型(前端渲染类型)
	Sort         interface{} // 排序值
	Description  interface{} // 字典描述
	CreateUserId interface{} // 创建用户ID
	CreateTime   *gtime.Time // 创建时间
	UpdateUserId interface{} // 修改用户ID
	UpdateTime   *gtime.Time // 修改时间
	Status       interface{} // 是否启用(0:禁用,1:启用)
	IsDeleted    interface{} // 是否删除(0:否,1:是)
}
