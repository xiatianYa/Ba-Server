package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" tags:"login" method:"post" summary:"用户登录(密码)"`
	Username string `json:"userName" v:"required|length:3,12"`
	Password string `json:"password" v:"required|length:6,16"`
}
type LoginRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type GetUserInfoReq struct {
	g.Meta `path:"/getUserInfo" tags:"getUserInfo" method:"get" summary:"获取用户信息"`
}
type GetUserInfoRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
