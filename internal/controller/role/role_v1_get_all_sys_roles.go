package role

import (
	"Ba-Server/internal/service"
	"context"

	"Ba-Server/api/role/v1"
)

func (c *ControllerV1) GetAllSysRoles(ctx context.Context, req *v1.GetAllSysRolesReq) (res *v1.GetAllSysRolesRes, err error) {
	out, err := service.Role().GetAllSysRolesReq(ctx)
	return (*v1.GetAllSysRolesRes)(&out), err
}
