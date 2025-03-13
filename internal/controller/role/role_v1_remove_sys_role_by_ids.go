package role

import (
	"Ba-Server/internal/service"
	"context"

	"Ba-Server/api/role/v1"
)

func (c *ControllerV1) RemoveSysRoleByIds(ctx context.Context, req *v1.RemoveSysRoleByIdsReq) (res *v1.RemoveSysRoleByIdsRes, err error) {
	out, err := service.Role().RemoveSysRoleByIds(ctx, req)
	return out, err
}
