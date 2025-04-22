package v1

import (
	"Ba-Server/internal/model/domain"
	"github.com/gogf/gf/v2/frame/g"
)

type GetSysUserPageReq struct {
	g.Meta     `path:"/page" tags:"page" method:"get" summary:"获取用户列表分页"`
	Current    int    `v:"required" json:"current"`
	Size       int    `v:"required" json:"size"`
	UserName   string `json:"userName"`
	UserGender int    `json:"userGender"`
	NickName   string `json:"nickName"`
	UserPhone  string `json:"userPhone"`
	UserEmail  string `json:"userEmail"`
	Status     string `json:"status"`
}

type GetSysUserPageRes *domain.RPage

type SaveSysUserReq struct {
	g.Meta     `path:"/save" tags:"page" method:"post" summary:"新增用户"`
	UserName   string   `v:"regex:^[\u4E00-\u9FA5a-zA-Z0-9_-]{4,16}$#用户名校验失败,请输入4-16位用户名"`
	NickName   string   `v:"regex:^[\u4E00-\u9FA5a-zA-Z0-9_-]{1,16}$#昵称校验失败,请输入1-16位昵称"`
	PassWord   string   `v:"regex:^\\w{6,18}$#密码校验失败,请输入6-18位密码"`
	UserGender int      `v:"required|in:1,2,3#用户性别为必填项,请输入有效的性别值（1 或 2 或 3）" json:"userGender"`
	UserPhone  string   `v:"regex:^1[3-9]\\d{9}$#手机号格式错误,请输入11位手机号" json:"userPhone"`
	UserEmail  string   `v:"email#请输入有效的邮箱地址" json:"userEmail"`
	Status     string   `v:"required" json:"status"`
	UserRoles  []string `v:"required" json:"userRoles"`
}

type SaveSysUserRes bool

type RemoveSysUserByIdsReq struct {
	g.Meta `path:"/removeByIds" tags:"removeByIds" method:"delete" summary:"删除多个用户"`
	Ids    []int64 `v:"required" json:"ids"`
}

type RemoveSysUserByIdsRes bool

type RemoveSysUserByIdReq struct {
	g.Meta `path:"/remove/{id}" tags:"remove" method:"delete" summary:"删除单个用户"`
	Id     int64 `v:"required" dc:"user id"`
}

type RemoveSysUserByIdRes bool

type UpdateSysUserReq struct {
	g.Meta     `path:"/update" tags:"update" method:"put" summary:"修改用户信息"`
	Id         int64    `v:"required" json:"id"`
	UserName   string   `v:"regex:^[\u4E00-\u9FA5a-zA-Z0-9_-]{4,16}$#用户名校验失败,请输入4-16位用户名"`
	NickName   string   `v:"regex:^[\u4E00-\u9FA5a-zA-Z0-9_-]{1,16}$#昵称校验失败,请输入1-16位昵称"`
	UserRoles  []string `v:"required" json:"userRoles"`
	UserGender int      `v:"required|in:1,2,3#用户性别为必填项,请输入有效的性别值（1 或 2 或 3）" json:"userGender"`
	Status     string   `v:"required" json:"status"`
	UserPhone  string   `v:"regex:^1[3-9]\\d{9}$#手机号格式错误,请输入11位手机号" json:"userPhone"`
	UserEmail  string   `v:"email#请输入有效的邮箱地址" json:"userEmail"`
}

type UpdateSysUserRes bool

type GetAllUserNameReq struct {
	g.Meta `path:"/allUserName" method:"get" summary:"获取所有用户枚举"`
}

type GetAllUserNameRes []domain.Options
