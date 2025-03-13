package role

import (
	"Ba-Server/internal/service"
	"context"

	"Ba-Server/api/role/v1"
)

func (c *ControllerV1) UpdateSysRole(ctx context.Context, req *v1.UpdateSysRoleReq) (res *v1.UpdateSysRoleRes, err error) {
	out, err := service.Role().UpdateSysRole(ctx, req)
	return out, err
}
