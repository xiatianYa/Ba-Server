package menu

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/menu/v1"
)

func (c *ControllerV1) SaveSysMenu(ctx context.Context, req *v1.SaveSysMenuReq) (res *v1.SaveSysMenuRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
