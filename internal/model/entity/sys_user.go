// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	Id            int64       `json:"id"            orm:"id"              ` // 主键ID
	UserName      string      `json:"userName"      orm:"user_name"       ` // 用户名称
	Password      string      `json:"password"      orm:"password"        ` // 密码
	NickName      string      `json:"nickName"      orm:"nick_name"       ` // 昵称
	UserEmail     string      `json:"userEmail"     orm:"user_email"      ` // 邮箱
	UserGender    int         `json:"userGender"    orm:"user_gender"     ` // 用户性别
	UserPhone     string      `json:"userPhone"     orm:"user_phone"      ` // 手机
	CreateTime    *gtime.Time `json:"createTime"    orm:"create_time"     ` // 创建时间
	UpdateTime    *gtime.Time `json:"updateTime"    orm:"update_time"     ` // 修改时间
	LastLoginTime *gtime.Time `json:"lastLoginTime" orm:"last_login_time" ` // 最后登录时间
	Salt          string      `json:"salt"          orm:"salt"            ` // MD5的盐值
	Status        string      `json:"status"        orm:"status"          ` // 是否启用(0:禁用,1:启用)
	IsDeleted     uint        `json:"isDeleted"     orm:"is_deleted"      ` // 是否删除(0:否,1:是)
}
