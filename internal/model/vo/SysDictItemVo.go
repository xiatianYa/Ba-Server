package vo

import "github.com/gogf/gf/v2/os/gtime"

// SysDictItemVo 定义子字典结构体
type SysDictItemVo struct {
	// 主键
	Id int64 `json:"id"`
	// 父字典 ID
	DictID int64 `json:"dictId" description:"父字典ID"`
	// 父字典编码
	DictCode string `json:"dictCode" description:"父字典编码"`
	// 数据值
	Value string `json:"value" description:"数据值"`
	// 中文名称
	ZhCn string `json:"zhCn" description:"中文名称"`
	// 英文名称
	EnUs string `json:"enUs" description:"英文名称"`
	// 类型(前端渲染类型)
	Type string `json:"type" description:"类型(前端渲染类型)"`
	// 排序值
	Sort int `json:"sort" description:"排序值"`
	// 字典描述
	Description string `json:"description" description:"字典描述"`
	// 是否启用(0:禁用,1:启用)
	Status     string      `json:"status" description:"是否启用(0:禁用,1:启用)"`
	CreateTime *gtime.Time `json:"createTime"` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime"` // 修改时间
}
