// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package role

import (
	"context"

	"Ba-Server/api/role/v1"
)

type IRoleV1 interface {
	GetAllSysRoles(ctx context.Context, req *v1.GetAllSysRolesReq) (res *v1.GetAllSysRolesRes, err error)
	GetSysRolePage(ctx context.Context, req *v1.GetSysRolePageReq) (res *v1.GetSysRolePageRes, err error)
	SaveSysRole(ctx context.Context, req *v1.SaveSysRoleReq) (res *v1.SaveSysRoleRes, err error)
	RemoveSysRoleByIds(ctx context.Context, req *v1.RemoveSysRoleByIdsReq) (res *v1.RemoveSysRoleByIdsRes, err error)
	RemoveSysRoleById(ctx context.Context, req *v1.RemoveSysRoleByIdReq) (res *v1.RemoveSysRoleByIdRes, err error)
	UpdateSysRole(ctx context.Context, req *v1.UpdateSysRoleReq) (res *v1.UpdateSysRoleRes, err error)
	UpdateRoleMenu(ctx context.Context, req *v1.UpdateRoleMenuReq) (res *v1.UpdateRoleMenuRes, err error)
	GetPermissionTree(ctx context.Context, req *v1.GetPermissionTreeReq) (res *v1.GetPermissionTreeRes, err error)
	GetPermissionIdsByRoleId(ctx context.Context, req *v1.GetPermissionIdsByRoleIdReq) (res *v1.GetPermissionIdsByRoleIdRes, err error)
	UpdatePermissionIdsByRoleId(ctx context.Context, req *v1.UpdatePermissionIdsByRoleIdReq) (res *v1.UpdatePermissionIdsByRoleIdRes, err error)
}
