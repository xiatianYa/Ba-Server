package vo

import "github.com/gogf/gf/v2/os/gtime"

// SysDictVo is the golang structure for table sys_dict.
type SysDictVo struct {
	// 主键
	Id          int64       `json:"id"`
	Name        string      `json:"name" description:"字典名称"`
	Code        string      `json:"code" description:"字典编码"`
	Type        string      `json:"type" description:"字典类型(1:系统字典,2:业务字典)"`
	Sort        int         `json:"sort" description:"排序值"`
	Description string      `json:"description" description:"字典描述"`
	Status      string      `json:"status" description:"是否启用(0:禁用,1:启用)"`
	CreateTime  *gtime.Time `json:"createTime"` // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"` // 修改时间
}
