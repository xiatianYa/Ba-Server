package user

import (
	"Ba-Server/internal/service"
	"context"

	"Ba-Server/api/user/v1"
)

func (c *ControllerV1) RemoveSysUserById(ctx context.Context, req *v1.RemoveSysUserByIdReq) (res *v1.RemoveSysUserByIdRes, err error) {
	out, err := service.User().RemoveSysUserById(ctx, req)
	return out, err
}
