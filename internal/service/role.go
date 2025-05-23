// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "Ba-Server/api/role/v1"
	"Ba-Server/internal/model/vo"
	"context"
)

type (
	IRole interface {
		GetAllSysRolesReq(ctx context.Context) (rolesVo []vo.SysRoleVo, err error)
		GetSysRolePage(ctx context.Context, req *v1.GetSysRolePageReq) (total int, records []vo.SysRoleVo, err error)
		SaveSysRole(ctx context.Context, req *v1.SaveSysRoleReq) (res *v1.SaveSysRoleRes, err error)
		RemoveSysRoleByIds(ctx context.Context, req *v1.RemoveSysRoleByIdsReq) (res *v1.RemoveSysRoleByIdsRes, err error)
		RemoveSysRoleById(ctx context.Context, req *v1.RemoveSysRoleByIdReq) (res *v1.RemoveSysRoleByIdRes, err error)
		UpdateSysRole(ctx context.Context, req *v1.UpdateSysRoleReq) (res *v1.UpdateSysRoleRes, err error)
		UpdateRoleMenu(ctx context.Context, req *v1.UpdateRoleMenuReq) (res *v1.UpdateRoleMenuRes, err error)
		GetPermissionTree(ctx context.Context, req *v1.GetPermissionTreeReq) (res []vo.SysPermissionTreeVo, err error)
		GetPermissionIdsByRoleId(ctx context.Context, req *v1.GetPermissionIdsByRoleIdReq) (res []int64, err error)
		UpdatePermissionIdsByRoleId(ctx context.Context, req *v1.UpdatePermissionIdsByRoleIdReq) (res *v1.UpdatePermissionIdsByRoleIdRes, err error)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
