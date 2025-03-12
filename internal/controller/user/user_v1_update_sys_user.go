package user

import (
	"Ba-Server/internal/service"
	"context"

	"Ba-Server/api/user/v1"
)

func (c *ControllerV1) UpdateSysUser(ctx context.Context, req *v1.UpdateSysUserReq) (res *v1.UpdateSysUserRes, err error) {
	out, err := service.User().UpdateSysUser(ctx, req)
	return out, err
}
