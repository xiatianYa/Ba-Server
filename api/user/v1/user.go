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
