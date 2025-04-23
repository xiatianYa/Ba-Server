package route

import (
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/route/v1"
)

func (c *ControllerV1) GetConstantRoutes(ctx context.Context, req *v1.GetConstantRoutesReq) (res *v1.GetConstantRoutesRes, err error) {
	out, err := service.Route().GetConstantRoutes(ctx)
	if err != nil {
		return nil, gerror.New("获取用户常量路由异常" + err.Error())
	}
	return (*v1.GetConstantRoutesRes)(&out), nil
}
