package menu

import (
	"Ba-Server/internal/service"
	"context"

	"Ba-Server/api/menu/v1"
)

func (c *ControllerV1) GetMenuIdsByRoleId(ctx context.Context, req *v1.GetMenuIdsByRoleIdReq) (res *v1.GetMenuIdsByRoleIdRes, err error) {
	out, err := service.Menu().GetMenuIdsByRoleId(ctx, req)
	if err != nil {
		return nil, err
	}
	return (*v1.GetMenuIdsByRoleIdRes)(&out), nil
}
