package vo

// SysUserInfoVo 用户信息 VO 展示类
type SysUserInfoVo struct {
	// 用户ID
	UserID int64 `json:"userId"`
	// 用户名称
	UserName string `json:"userName"`
	// 用户手机号
	UserPhone string `json:"userPhone"`
	// 用户邮箱
	UserEmail string `json:"userEmail"`
	// 用户性别
	UserGender int `json:"userGender"`
	// 用户权限列表
	RoleCodes []string `json:"roles"`
	// 用户权限按钮列表
	Buttons []string `json:"buttons"`
}
