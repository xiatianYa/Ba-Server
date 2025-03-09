package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetUserInfoReq struct {
	g.Meta     `path:"/getUserInfo" tags:"getUserInfo" method:"get" summary:"获取用户信息"`
	Permission string `permission:"admin_permission"`
}
type GetUserInfoRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
