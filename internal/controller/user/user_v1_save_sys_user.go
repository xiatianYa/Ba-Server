package user

import (
	"Ba-Server/internal/service"
	"context"

	"Ba-Server/api/user/v1"
)

func (c *ControllerV1) SaveSysUser(ctx context.Context, req *v1.SaveSysUserReq) (res *v1.SaveSysUserRes, err error) {
	out, err := service.User().SaveSysUser(ctx, req)
	return out, err
}
