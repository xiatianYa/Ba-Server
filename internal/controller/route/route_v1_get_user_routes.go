package route

import (
	"Ba-Server/api/route/v1"
	"Ba-Server/internal/service"
	"context"
)

func (c *ControllerV1) GetUserRoutes(ctx context.Context, req *v1.GetUserRoutesReq) (res *v1.GetUserRoutesRes, err error) {
	out, err := service.Route().GetUserRoutes(ctx)
	if err != nil {
		return nil, err
	}
	return (*v1.GetUserRoutesRes)(&out), nil
}
