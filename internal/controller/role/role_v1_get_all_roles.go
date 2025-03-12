package role

import (
	"Ba-Server/internal/service"
	"context"

	"Ba-Server/api/role/v1"
)

func (c *ControllerV1) GetAllRoles(ctx context.Context, req *v1.GetAllRolesReq) (res *v1.GetAllRolesRes, err error) {
	out, err := service.Role().GetAllRolesReq(ctx)
	if err != nil {
		return nil, err
	}
	return (*v1.GetAllRolesRes)(&out), nil
}
