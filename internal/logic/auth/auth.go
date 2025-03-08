package auth

import (
	v1 "Ba-Server/api/auth/v1"
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"
)

type (
	sAuth struct{}
)

func init() {
	service.RegisterAuth(New())
}

func New() service.IAuth {
	return &sAuth{}
}

func (s sAuth) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (constantRoutes vo.SysUserInfoVo, err error) {
	return
}
