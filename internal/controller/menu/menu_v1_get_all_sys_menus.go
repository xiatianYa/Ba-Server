package menu

import (
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/menu/v1"
)

func (c *ControllerV1) GetAllSysMenus(ctx context.Context, req *v1.GetAllSysMenusReq) (res *v1.GetAllSysMenusRes, err error) {
	out, err := service.Menu().GetAllSysMenus(ctx, req)
	if err != nil {
		return nil, gerror.New("获取菜单列表失败")
	}
	return (*v1.GetAllSysMenusRes)(&out), nil
}
