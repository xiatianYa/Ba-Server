package menu

import (
	"Ba-Server/internal/service"
	"context"

	"Ba-Server/api/menu/v1"
)

func (c *ControllerV1) GetAllSysMenus(ctx context.Context, req *v1.GetAllSysMenusReq) (res *v1.GetAllSysMenusRes, err error) {
	out, err := service.Menu().GetAllSysMenus(ctx, req)
	if err != nil {
		return nil, err
	}
	return (*v1.GetAllSysMenusRes)(&out), nil
}
