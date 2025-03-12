// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package role

import (
	"context"

	"Ba-Server/api/role/v1"
)

type IRoleV1 interface {
	GetAllRoles(ctx context.Context, req *v1.GetAllRolesReq) (res *v1.GetAllRolesRes, err error)
}
