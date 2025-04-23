package user

import (
	"Ba-Server/api/user/v1"
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) RemoveSysUserByIds(ctx context.Context, req *v1.RemoveSysUserByIdsReq) (res *v1.RemoveSysUserByIdsRes, err error) {
	out, err := service.User().RemoveSysUserByIds(ctx, req)
	if err != nil {
		return nil, gerror.New("删除用户失败" + err.Error())
	}
	return out, err
}
