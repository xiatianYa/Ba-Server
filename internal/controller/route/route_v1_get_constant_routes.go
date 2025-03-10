package route

import (
	"Ba-Server/internal/service"
	"context"

	"Ba-Server/api/route/v1"
)

func (c *ControllerV1) GetConstantRoutes(ctx context.Context, req *v1.GetConstantRoutesReq) (res *v1.GetConstantRoutesRes, err error) {
	out, err := service.Route().GetConstantRoutes(ctx)
	if err != nil {
		return nil, err
	}
	return (*v1.GetConstantRoutesRes)(&out), nil
}
