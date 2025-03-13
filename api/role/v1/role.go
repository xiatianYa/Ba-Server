package v1

import (
	"Ba-Server/internal/model/domain"
	"Ba-Server/internal/model/vo"
	"github.com/gogf/gf/v2/frame/g"
)

type GetAllSysRolesReq struct {
	g.Meta `path:"/getAllRoles" tags:"getAllRoles" method:"get" summary:"获取权限标识列表"`
}

type GetAllSysRolesRes []vo.SysRoleVo

type GetSysRolePageReq struct {
	g.Meta   `path:"/page" tags:"page" method:"get" summary:"获取权限列表分页"`
	Current  int    `v:"required" json:"current"`
	Size     int    `v:"required" json:"size"`
	RoleName string `json:"roleName"`
	RoleCode string `json:"roleCode"`
	Status   string `json:"status"`
}

type GetSysRolePageRes *domain.RPage

type SaveSysRoleReq struct {
	g.Meta   `path:"/save" tags:"save" method:"post" summary:"新增角色列表"`
	RoleName string `v:"required" json:"roleName"`
	RoleCode string `v:"required" json:"roleCode"`
	RoleDesc string `json:"roleDesc"`
	Status   string `v:"required" json:"status"`
}

type SaveSysRoleRes bool

type RemoveSysRoleByIdsReq struct {
	g.Meta `path:"/removeByIds" tags:"removeByIds" method:"delete" summary:"删除多个角色"`
	Ids    []int64 `v:"required" json:"ids"`
}

type RemoveSysRoleByIdsRes bool

type RemoveSysRoleByIdReq struct {
	g.Meta `path:"/remove/{id}" tags:"remove" method:"delete" summary:"删除单个角色"`
	Id     int64 `v:"required" dc:"role id"`
}

type RemoveSysRoleByIdRes bool

type UpdateSysRoleReq struct {
	g.Meta   `path:"/update" tags:"update" method:"put" summary:"修改角色信息"`
	Id       int64  `v:"required" json:"id"`
	RoleName string `v:"required" json:"roleName"`
	RoleCode string `v:"required" json:"roleCode"`
	RoleDesc string `json:"roleDesc"`
	Status   string `v:"required" json:"status"`
}

type UpdateSysRoleRes bool

type UpdateRoleMenuReq struct {
	g.Meta  `path:"/updateRoleMenu" tags:"updateRoleMenu" method:"put" summary:"修改角色菜单权限"`
	RoleId  int64   `v:"required" json:"roleId"`
	MenuIds []int64 `json:"menuIds"`
}

type UpdateRoleMenuRes bool
