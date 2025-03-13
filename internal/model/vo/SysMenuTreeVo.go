package vo

// SysMenuTreeVo is the golang structure for table sys_menu.
type SysMenuTreeVo struct {
	Id       int64           `json:"id"`       // 主键
	Label    string          `json:"label"`    // 路由名称
	Pid      string          `json:"pid"`      // 父路由ID
	Sort     string          `json:"sort"`     // 排序
	Children []SysMenuTreeVo `json:"children"` //字路由
}
