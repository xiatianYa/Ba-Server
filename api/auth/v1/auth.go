package v1

import (
	"Ba-Server/internal/model/vo"
	"github.com/gogf/gf/v2/frame/g"
)

type GetUserInfoReq struct {
	g.Meta `path:"/getUserInfo" tags:"getUserInfo" method:"get" summary:"获取用户信息"`
}
type GetUserInfoRes *vo.SysUserInfoVo
