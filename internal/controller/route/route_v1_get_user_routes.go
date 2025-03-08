package route

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/route/v1"
)

func (c *ControllerV1) GetUserRoutes(ctx context.Context, req *v1.GetUserRoutesReq) (res *v1.GetUserRoutesRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
