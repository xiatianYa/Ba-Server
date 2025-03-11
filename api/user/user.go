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
}
