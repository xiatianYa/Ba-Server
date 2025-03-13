package user

import (
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/user/v1"
)

func (c *ControllerV1) UpdateSysUser(ctx context.Context, req *v1.UpdateSysUserReq) (res *v1.UpdateSysUserRes, err error) {
	out, err := service.User().UpdateSysUser(ctx, req)
	if err != nil {
		return nil, gerror.New("修改用户失败")
	}
	return out, err
}
