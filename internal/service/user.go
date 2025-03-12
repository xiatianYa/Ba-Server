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
		GetSysUserPage(ctx context.Context, req *v1.GetSysUserPageReq) (total int, records []*vo.SysUserVo, err error)
		GetUserInfoVo(ctx context.Context, sysUser entity.SysUser) (sysUserInfoVo *vo.SysUserInfoVo, err error)
		SaveSysUser(ctx context.Context, req *v1.SaveSysUserReq) (res *v1.SaveSysUserRes, err error)
		RemoveSysUserByIds(ctx context.Context, req *v1.RemoveSysUserByIdsReq) (res *v1.RemoveSysUserByIdsRes, err error)
		RemoveSysUserById(ctx context.Context, req *v1.RemoveSysUserByIdReq) (res *v1.RemoveSysUserByIdRes, err error)
		UpdateSysUser(ctx context.Context, req *v1.UpdateSysUserReq) (res *v1.UpdateSysUserRes, err error)
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
