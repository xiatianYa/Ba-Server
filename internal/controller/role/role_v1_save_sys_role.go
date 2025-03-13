package role

import (
	"Ba-Server/internal/service"
	"context"

	"Ba-Server/api/role/v1"
)

func (c *ControllerV1) SaveSysRole(ctx context.Context, req *v1.SaveSysRoleReq) (res *v1.SaveSysRoleRes, err error) {
	out, err := service.Role().SaveSysRole(ctx, req)
	return out, err
}
