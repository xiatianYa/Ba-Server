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
	UserName   string   `v:"required" json:"userName"`
	PassWord   string   `v:"required" json:"passWord"`
	UserGender int      `v:"required" json:"userGender"`
	NickName   string   `v:"required" json:"nickName"`
	UserPhone  string   `json:"userPhone"`
	UserEmail  string   `json:"userEmail"`
	Status     string   `v:"required" json:"status"`
	UserRoles  []string `v:"required" json:"userRoles"`
}

type SaveSysUserRes bool

type RemoveSysUserByIdsReq struct {
	g.Meta `path:"/removeByIds" tags:"removeByIds" method:"delete" summary:"删除多个用户"`
	Ids    []int64 `v:"required" json:"ids"`
}

type RemoveSysUserByIdsRes bool
