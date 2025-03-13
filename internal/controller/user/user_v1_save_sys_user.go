package user

import (
	"Ba-Server/api/user/v1"
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) SaveSysUser(ctx context.Context, req *v1.SaveSysUserReq) (res *v1.SaveSysUserRes, err error) {
	out, err := service.User().SaveSysUser(ctx, req)
	if err != nil {
		return nil, gerror.New("新增用户失败")
	}
	return out, err
}
