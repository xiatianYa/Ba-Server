package vo

// SysUserVo is the golang structure for table sys_user.
type SysUserVo struct {
	Id         int64    `json:"id"`         // 主键ID
	UserName   string   `json:"userName"`   // 用户名称
	NickName   string   `json:"nickName"`   // 昵称
	Avatar     string   `json:"avatar"`     // 头像
	UserEmail  string   `json:"userEmail"`  // 邮箱
	UserGender string   `json:"userGender"` // 用户性别
	UserPhone  string   `json:"userPhone"`  // 手机
	Status     string   `json:"status"`     // 是否启用(0:禁用,1:启用)
	UserRoles  []string `json:"userRoles"`  // 用户角色标识列表
}
