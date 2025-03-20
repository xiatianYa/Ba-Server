// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package menu

import (
	"context"

	"Ba-Server/api/menu/v1"
)

type IMenuV1 interface {
	GetAllSysMenus(ctx context.Context, req *v1.GetAllSysMenusReq) (res *v1.GetAllSysMenusRes, err error)
	GetSysMenuTree(ctx context.Context, req *v1.GetSysMenuTreeReq) (res *v1.GetSysMenuTreeRes, err error)
	GetSysMenuIdsByRoleId(ctx context.Context, req *v1.GetSysMenuIdsByRoleIdReq) (res *v1.GetSysMenuIdsByRoleIdRes, err error)
	GetSysMenuPage(ctx context.Context, req *v1.GetSysMenuPageReq) (res *v1.GetSysMenuPageRes, err error)
	SaveSysMenu(ctx context.Context, req *v1.SaveSysMenuReq) (res *v1.SaveSysMenuRes, err error)
	RemoveSysMenuByIds(ctx context.Context, req *v1.RemoveSysMenuByIdsReq) (res *v1.RemoveSysMenuByIdsRes, err error)
	RemoveSysMenuById(ctx context.Context, req *v1.RemoveSysMenuByIdReq) (res *v1.RemoveSysMenuByIdRes, err error)
	UpdateSysMenu(ctx context.Context, req *v1.UpdateSysMenuReq) (res *v1.UpdateSysMenuRes, err error)
}
