// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictItem is the golang structure for table sys_dict_item.
type SysDictItem struct {
	Id          int64       `json:"id"          orm:"id"          ` // 主键
	DictId      int64       `json:"dictId"      orm:"dict_id"     ` // 父字典ID
	DictCode    string      `json:"dictCode"    orm:"dict_code"   ` // 父字典编码
	Value       string      `json:"value"       orm:"value"       ` // 数据值
	ZhCn        string      `json:"zhCn"        orm:"zh_cn"       ` // 中文名称
	EnUs        string      `json:"enUs"        orm:"en_us"       ` // 英文名称
	Type        string      `json:"type"        orm:"type"        ` // 类型(前端渲染类型)
	Sort        int         `json:"sort"        orm:"sort"        ` // 排序值
	Description string      `json:"description" orm:"description" ` // 字典描述
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time" ` // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time" ` // 修改时间
	Status      string      `json:"status"      orm:"status"      ` // 是否启用(0:禁用,1:启用)
	IsDeleted   uint        `json:"isDeleted"   orm:"is_deleted"  ` // 是否删除(0:否,1:是)
}
