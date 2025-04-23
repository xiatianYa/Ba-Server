package route

import (
	"Ba-Server/api/route/v1"
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) GetUserRoutes(ctx context.Context, req *v1.GetUserRoutesReq) (res *v1.GetUserRoutesRes, err error) {
	out, err := service.Route().GetUserRoutes(ctx)
	if err != nil {
		return nil, gerror.New("获取用户动态路由异常" + err.Error())
	}
	return (*v1.GetUserRoutesRes)(&out), nil
}
