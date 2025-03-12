// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"Ba-Server/api/user/v1"
)

type IUserV1 interface {
	GetSysUserPage(ctx context.Context, req *v1.GetSysUserPageReq) (res *v1.GetSysUserPageRes, err error)
	SaveSysUser(ctx context.Context, req *v1.SaveSysUserReq) (res *v1.SaveSysUserRes, err error)
	RemoveSysUserByIds(ctx context.Context, req *v1.RemoveSysUserByIdsReq) (res *v1.RemoveSysUserByIdsRes, err error)
	RemoveSysUserById(ctx context.Context, req *v1.RemoveSysUserByIdReq) (res *v1.RemoveSysUserByIdRes, err error)
	UpdateSysUser(ctx context.Context, req *v1.UpdateSysUserReq) (res *v1.UpdateSysUserRes, err error)
}
