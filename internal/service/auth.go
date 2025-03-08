// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "Ba-Server/api/auth/v1"
	"Ba-Server/internal/model/vo"
	"context"
)

type (
	IAuth interface {
		UserNameLogin(ctx context.Context, req *v1.LoginReq) (constantRoutes map[string]interface{}, err error)
		GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (constantRoutes vo.SysUserInfoVo, err error)
	}
)

var (
	localAuth IAuth
)

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
