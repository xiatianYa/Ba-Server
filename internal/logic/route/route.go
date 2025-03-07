package route

import (
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"
)

type (
	sRoute struct{}
)

func init() {
	service.RegisterRoute(New())
}

func New() service.IRoute {
	return &sRoute{}
}

func (s sRoute) GetConstantRoutes(ctx context.Context) (constantRoutes map[string]string, err error) {
	//TODO implement me
	panic("implement me")
}

func (s sRoute) GetUserRoutes(ctx context.Context) (userRouteVo vo.SysUserRouteVO, err error) {
	//TODO implement me
	panic("implement me")
}
