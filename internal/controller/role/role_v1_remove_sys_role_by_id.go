package role

import (
	"Ba-Server/internal/service"
	"context"

	"Ba-Server/api/role/v1"
)

func (c *ControllerV1) RemoveSysRoleById(ctx context.Context, req *v1.RemoveSysRoleByIdReq) (res *v1.RemoveSysRoleByIdRes, err error) {
	out, err := service.Role().RemoveSysRoleById(ctx, req)
	return out, err
}
