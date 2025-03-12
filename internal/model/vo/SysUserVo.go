package vo

import "github.com/gogf/gf/v2/os/gtime"

// SysUserVo is the golang structure for table sys_user.
type SysUserVo struct {
	Id                 int64       `json:"id"                 orm:"id"                   ` // 主键ID
	UserName           string      `json:"userName"           orm:"user_name"            ` // 用户名称
	NickName           string      `json:"nickName"           orm:"nick_name"            ` // 昵称
	Avatar             string      `json:"avatar"             orm:"avatar"               ` // 头像
	UserEmail          string      `json:"userEmail"          orm:"user_email"           ` // 邮箱
	UserGender         string      `json:"userGender"         orm:"user_gender"          ` // 用户性别
	UserPhone          string      `json:"userPhone"          orm:"user_phone"           ` // 手机
	CreateUserId       int64       `json:"createUserId"       orm:"create_user_id"       ` // 创建用户ID
	CreateTime         *gtime.Time `json:"createTime"         orm:"create_time"          ` // 创建时间
	UpdateUserId       int64       `json:"updateUserId"       orm:"update_user_id"       ` // 修改用户ID
	UpdateTime         *gtime.Time `json:"updateTime"         orm:"update_time"          ` // 修改时间
	LastLoginTime      *gtime.Time `json:"lastLoginTime"      orm:"last_login_time"      ` // 最后登录时间
	UpdatePasswordTime *gtime.Time `json:"updatePasswordTime" orm:"update_password_time" ` // 修改密码时间
	Status             string      `json:"status"             orm:"status"               ` // 是否启用(0:禁用,1:启用)
	IsDeleted          uint        `json:"isDeleted"          orm:"is_deleted"           ` // 是否删除(0:否,1:是)
	UserRoles          []string    `json:"userRoles"          orm:"user_roles"           ` // 用户角色标识列表
}
