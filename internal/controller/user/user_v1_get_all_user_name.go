package user

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/user/v1"
)

func (c *ControllerV1) GetAllUserName(ctx context.Context, req *v1.GetAllUserNameReq) (res *v1.GetAllUserNameRes, err error) {
	out, err := service.User().GetAllUserName(ctx)
	if err != nil {
		return nil, gerror.New("获取用户名称配置异常" + err.Error())
	}
	return (*v1.GetAllUserNameRes)(&out), err
}
