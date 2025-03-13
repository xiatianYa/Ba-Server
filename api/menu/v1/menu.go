package v1

import (
	"Ba-Server/internal/model/vo"
	"github.com/gogf/gf/v2/frame/g"
)

type GetAllSysMenusReq struct {
	g.Meta `path:"/getAllMenus" tags:"getAllMenus" method:"get" summary:"获取菜单名称列表"`
}

type GetAllSysMenusRes []string

type GetSysMenuTreeReq struct {
	g.Meta `path:"/getMenuTree" tags:"getMenuTree" method:"get" summary:"获取菜单树"`
}

type GetSysMenuTreeRes []vo.SysMenuTreeVo

type GetMenuIdsByRoleIdReq struct {
	g.Meta `path:"/getMenuIdsByRoleId/{id}" tags:"getMenuIdsByRoleId" method:"get" summary:"获取菜单Ids根据角色Id"`
	Id     int64 `v:"required" dc:"role id"`
}

type GetMenuIdsByRoleIdRes []int64
