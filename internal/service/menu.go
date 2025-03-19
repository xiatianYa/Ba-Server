// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "Ba-Server/api/menu/v1"
	"Ba-Server/internal/model/vo"
	"context"
)

type (
	IMenu interface {
		GetAllSysMenus(ctx context.Context, req *v1.GetAllSysMenusReq) (res []string, err error)
		GetSysMenuTree(ctx context.Context, req *v1.GetSysMenuTreeReq) (res []vo.SysMenuTreeVo, err error)
		GetSysMenuIdsByRoleId(ctx context.Context, req *v1.GetSysMenuIdsByRoleIdReq) (res []int64, err error)
		GetSysMenuPage(ctx context.Context, req *v1.GetSysMenuPageReq) (total int, records []vo.SysMenuVo, err error)
		SaveSysMenu(ctx context.Context, req *v1.SaveSysMenuReq) (res *v1.SaveSysMenuRes, err error)
	}
)

var (
	localMenu IMenu
)

func Menu() IMenu {
	if localMenu == nil {
		panic("implement not found for interface IMenu, forgot register?")
	}
	return localMenu
}

func RegisterMenu(i IMenu) {
	localMenu = i
}
