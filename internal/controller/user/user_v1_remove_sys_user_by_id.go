package user

import (
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/user/v1"
)

func (c *ControllerV1) RemoveSysUserById(ctx context.Context, req *v1.RemoveSysUserByIdReq) (res *v1.RemoveSysUserByIdRes, err error) {
	out, err := service.User().RemoveSysUserById(ctx, req)
	//发送错误
	if err != nil {
		return nil, gerror.New("删除用户失败" + err.Error())
	}
	return out, err
}
