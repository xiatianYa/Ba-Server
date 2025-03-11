package v1

import (
	"Ba-Server/internal/model/vo"
	"github.com/gogf/gf/v2/frame/g"
)

type GetConstantRoutesReq struct {
	g.Meta `path:"/getConstantRoutes" tags:"getConstantRoutes" method:"get" summary:"获取用户常量路由"`
}

type GetConstantRoutesRes []map[string]interface{}

type GetUserRoutesReq struct {
	g.Meta `path:"/getUserRoutes" tags:"getUserRoutes" method:"get" summary:"获取用户权限路由"`
}
type GetUserRoutesRes *vo.SysUserRouteVO
