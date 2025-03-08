package route

import (
	"Ba-Server/internal/model/api"
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"Ba-Server/api/route/v1"
)

func (c *ControllerV1) GetUserRoutes(ctx context.Context, req *v1.GetUserRoutesReq) (res *v1.GetUserRoutesRes, err error) {
	routes, err := service.Route().GetUserRoutes(ctx)
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.WriteJson(api.Result{Code: 200, Data: routes, Msg: "success"})
	return
}
