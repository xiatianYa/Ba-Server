package vo

// SysDictItemOptionsVo is the golang structure for table sys_menu.
type SysDictItemOptionsVo struct {
	Label string `json:"label"` // 名称
	Value string `json:"value"` // 值
	Sort  int    `json:"sort"`  // 排序
	Type  string `json:"type"`  //类型(前端渲染类型)
}
