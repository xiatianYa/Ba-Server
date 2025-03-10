package service

import (
	"Ba-Server/internal/model/vo"
	"context"
)

type (
	IRoute interface {
		// GetConstantRoutes 获取常量路由
		GetConstantRoutes(ctx context.Context) (constantRoutes []map[string]interface{}, err error)
		// GetUserRoutes 获取用户路由
		GetUserRoutes(ctx context.Context) (userRouteVo vo.SysUserRouteVO, err error)
	}
)

var (
	localRoute IRoute
)

func Route() IRoute {
	if localRoute == nil {
		panic("implement not found for interface ICompany, forgot register?")
	}
	return localRoute
}

func RegisterRoute(i IRoute) {
	localRoute = i
}
