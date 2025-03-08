package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetConstantRoutesReq struct {
	g.Meta `path:"/getConstantRoutes" tags:"getConstantRoutes" method:"get" summary:"获取用户常量录音"`
}

type GetConstantRoutesRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type GetUserRoutesReq struct {
	g.Meta `path:"/getUserRoutes" tags:"getUserRoutes" method:"get" summary:"获取用户权限路由"`
}
type GetUserRoutesRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
