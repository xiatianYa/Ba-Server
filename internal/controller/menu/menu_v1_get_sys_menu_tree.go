package menu

import (
	"Ba-Server/internal/service"
	"context"

	"Ba-Server/api/menu/v1"
)

func (c *ControllerV1) GetSysMenuTree(ctx context.Context, req *v1.GetSysMenuTreeReq) (res *v1.GetSysMenuTreeRes, err error) {
	out, err := service.Menu().GetSysMenuTree(ctx, req)
	if err != nil {
		return nil, err
	}
	return (*v1.GetSysMenuTreeRes)(&out), nil
}
