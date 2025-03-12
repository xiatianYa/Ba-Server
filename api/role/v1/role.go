package v1

import (
	"Ba-Server/internal/model/vo"
	"github.com/gogf/gf/v2/frame/g"
)

type GetAllRolesReq struct {
	g.Meta `path:"/getAllRoles" tags:"getAllRoles" method:"get" summary:"获取权限标识列表"`
}

type GetAllRolesRes []vo.SysRoleVo
