package menu

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/menu/v1"
)

func (c *ControllerV1) SaveSysMenu(ctx context.Context, req *v1.SaveSysMenuReq) (res *v1.SaveSysMenuRes, err error) {
	out, err := service.Menu().SaveSysMenu(ctx, req)
	if err != nil {
		return nil, gerror.New("新增菜单失败" + err.Error())
	}
	return out, err
}
