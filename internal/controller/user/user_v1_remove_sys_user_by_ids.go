package user

import (
	"Ba-Server/api/user/v1"
	"Ba-Server/internal/service"
	"context"
)

func (c *ControllerV1) RemoveSysUserByIds(ctx context.Context, req *v1.RemoveSysUserByIdsReq) (res *v1.RemoveSysUserByIdsRes, err error) {
	out, err := service.User().RemoveSysUserByIds(ctx, req)
	return out, err
}
