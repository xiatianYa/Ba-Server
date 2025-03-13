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
		GetMenuIdsByRoleId(ctx context.Context, req *v1.GetMenuIdsByRoleIdReq) (res []int64, err error)
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
