package vo

import "github.com/gogf/gf/v2/os/gtime"

// SysUserVo is the golang structure for table sys_user.
type SysUserVo struct {
	Id                 int64       `json:"id"`                 // 主键ID
	UserName           string      `json:"userName"`           // 用户名称
	NickName           string      `json:"nickName"`           // 昵称
	Avatar             string      `json:"avatar"`             // 头像
	UserEmail          string      `json:"userEmail"`          // 邮箱
	UserGender         string      `json:"userGender"`         // 用户性别
	UserPhone          string      `json:"userPhone"`          // 手机
	CreateUserId       int64       `json:"createUserId"`       // 创建用户ID
	CreateTime         *gtime.Time `json:"createTime"`         // 创建时间
	UpdateUserId       int64       `json:"updateUserId"`       // 修改用户ID
	UpdateTime         *gtime.Time `json:"updateTime" `        // 修改时间
	LastLoginTime      *gtime.Time `json:"lastLoginTime"`      // 最后登录时间
	UpdatePasswordTime *gtime.Time `json:"updatePasswordTime"` // 修改密码时间
	Status             string      `json:"status"`             // 是否启用(0:禁用,1:启用)
	IsDeleted          uint        `json:"isDeleted"`          // 是否删除(0:否,1:是)
	UserRoles          []string    `json:"userRoles"`          // 用户角色标识列表
}
