package menu

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/menu/v1"
)

func (c *ControllerV1) RemoveSysMenuByIds(ctx context.Context, req *v1.RemoveSysMenuByIdsReq) (res *v1.RemoveSysMenuByIdsRes, err error) {
	out, err := service.Menu().RemoveSysMenuByIds(ctx, req)
	if err != nil {
		return nil, gerror.New("删除多个菜单失败" + err.Error())
	}
	return out, err
}
