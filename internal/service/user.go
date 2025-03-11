// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "Ba-Server/api/user/v1"
	"Ba-Server/internal/model/entity"
	"Ba-Server/internal/model/vo"
	"context"
)

type (
	IUser interface {
		GetSysUserPage(ctx context.Context, req *v1.GetSysUserPageReq) (pages int, total int, records []vo.SysUserVo, err error)
		GetUserInfoVo(ctx context.Context, sysUser entity.SysUser) (sysUserInfoVo *vo.SysUserInfoVo, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
