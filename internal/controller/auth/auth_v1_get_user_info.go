package auth

import (
	"Ba-Server/api/auth/v1"
	"Ba-Server/internal/service"
	"context"
)

func (c *ControllerV1) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (res *v1.GetUserInfoRes, err error) {
	out, err := service.Auth().GetUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return (*v1.GetUserInfoRes)(&out), nil
}
