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
	GetMenuIdsByRoleId(ctx context.Context, req *v1.GetMenuIdsByRoleIdReq) (res *v1.GetMenuIdsByRoleIdRes, err error)
}
