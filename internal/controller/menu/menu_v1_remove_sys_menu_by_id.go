package menu

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/menu/v1"
)

func (c *ControllerV1) RemoveSysMenuById(ctx context.Context, req *v1.RemoveSysMenuByIdReq) (res *v1.RemoveSysMenuByIdRes, err error) {
	out, err := service.Menu().RemoveSysMenuById(ctx, req)
	if err != nil {
		return nil, gerror.New("删除菜单失败" + err.Error())
	}
	return out, err
}
