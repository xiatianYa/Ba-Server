package menu

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/menu/v1"
)

func (c *ControllerV1) UpdateSysMenu(ctx context.Context, req *v1.UpdateSysMenuReq) (res *v1.UpdateSysMenuRes, err error) {
	out, err := service.Menu().UpdateSysMenu(ctx, req)
	if err != nil {
		return nil, gerror.New("修改菜单失败")
	}
	return out, err
}
