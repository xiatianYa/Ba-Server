package route

import (
	"Ba-Server/api/route/v1"
	"Ba-Server/internal/service"
	"context"
)

func (c *ControllerV1) GetConstantRoutes(ctx context.Context, req *v1.GetConstantRoutesReq) (res *v1.GetConstantRoutesRes, err error) {
	routes, err := service.Route().GetConstantRoutes(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.GetConstantRoutesRes{ConstantRoutes: routes}, nil
}
