package role

import (
	v1 "Ba-Server/api/role/v1"
	"Ba-Server/internal/consts"
	"Ba-Server/internal/dao"
	"Ba-Server/internal/model/do"
	"Ba-Server/internal/model/entity"
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"math/rand/v2"
)

type (
	sRole struct{}
)

func init() {
	service.RegisterRole(New())
}

func New() service.IRole {
	return &sRole{}
}

func (s sRole) GetAllSysRolesReq(ctx context.Context) (rolesVo []vo.SysRoleVo, err error) {
	var sysRoleVos []vo.SysRoleVo
	roleModel := dao.SysRole.Ctx(ctx)
	_ = roleModel.Where("status = ?", consts.ONE).Scan(&sysRoleVos)
	return sysRoleVos, nil
}

func (s sRole) GetSysRolePage(ctx context.Context, req *v1.GetSysRolePageReq) (total int, records []*vo.SysRoleVo, err error) {
	// 创建指针切片
	var result []*vo.SysRoleVo

	roleModel := dao.SysRole.Ctx(ctx)
	pageQuery := roleModel.Page(req.Current, req.Size)
	// 处理各个查询条件，只添加非空的条件
	if req.RoleName != "" {
		pageQuery = pageQuery.Where("role_name like ?", "%"+req.RoleName+"%")
	}
	if req.RoleCode != "" {
		pageQuery = pageQuery.Where("role_code like ?", "%"+req.RoleCode+"%")
	}
	if req.Status != "" {
		pageQuery = pageQuery.Where("status = ?", req.Status)
	}

	err = pageQuery.ScanAndCount(&result, &total, true)

	if err != nil {
		return 0, nil, err
	}

	// 如果列表为空
	if result == nil {
		records = []*vo.SysRoleVo{}
		return
	}
	records = result
	return
}

func (s sRole) SaveSysRole(ctx context.Context, req *v1.SaveSysRoleReq) (res *v1.SaveSysRoleRes, err error) {
	roleModel := dao.SysRole.Ctx(ctx)
	sysRole := do.SysRole{
		RoleName: req.RoleName,
		RoleCode: req.RoleCode,
		RoleDesc: req.RoleDesc,
		Status:   req.Status,
	}
	_, err = roleModel.Insert(sysRole)
	if err != nil {
		return nil, err
	}
	return
}

func (s sRole) RemoveSysRoleByIds(ctx context.Context, req *v1.RemoveSysRoleByIdsReq) (res *v1.RemoveSysRoleByIdsRes, err error) {
	if req.Ids == nil {
		return nil, gerror.New("用户IDS为空")
	}
	roleModel := dao.SysRole.Ctx(ctx)
	userRoleModel := dao.SysUserRole.Ctx(ctx)
	roleMenuModel := dao.SysRoleMenu.Ctx(ctx)
	rolePermissionModel := dao.SysRolePermission.Ctx(ctx)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		//先删除用户角色对应表
		_, err2 := userRoleModel.WhereIn("role_id", req.Ids).Delete()
		if err2 != nil {
			return err2
		}
		//删除角色权限对应列表
		_, err3 := rolePermissionModel.WhereIn("role_id", req.Ids).Delete()
		if err3 != nil {
			return err3
		}
		//删除当前角色
		_, err4 := roleModel.WhereIn("id", req.Ids).Delete()
		if err4 != nil {
			return err4
		}
		//删除角色菜单对应列表
		_, err5 := roleMenuModel.Where("role_id", req.Ids).Delete()
		if err5 != nil {
			return err5
		}
		return nil
	})
	if err != nil {
		return nil, gerror.New("角色删除失败")
	}
	return
}

func (s sRole) RemoveSysRoleById(ctx context.Context, req *v1.RemoveSysRoleByIdReq) (res *v1.RemoveSysRoleByIdRes, err error) {
	if req.Id == 0 {
		return nil, gerror.New("用户ID为空")
	}
	roleModel := dao.SysRole.Ctx(ctx)
	userRoleModel := dao.SysUserRole.Ctx(ctx)
	roleMenuModel := dao.SysRoleMenu.Ctx(ctx)
	rolePermissionModel := dao.SysRolePermission.Ctx(ctx)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		//先删除用户角色对应表
		_, err2 := userRoleModel.Where("role_id", req.Id).Delete()
		if err2 != nil {
			return err2
		}
		//删除角色权限对应列表
		_, err3 := rolePermissionModel.Where("role_id", req.Id).Delete()
		if err3 != nil {
			return err3
		}
		//删除当前角色
		_, err4 := roleModel.Where("id", req.Id).Delete()
		if err4 != nil {
			return err4
		}
		//删除角色菜单对应列表
		_, err5 := roleMenuModel.Where("role_id", req.Id).Delete()
		if err5 != nil {
			return err5
		}
		return nil
	})
	if err != nil {
		return nil, gerror.New("角色删除失败")
	}
	return
}

func (s sRole) UpdateSysRole(ctx context.Context, req *v1.UpdateSysRoleReq) (res *v1.UpdateSysRoleRes, err error) {
	roleModel := dao.SysRole.Ctx(ctx)
	sysRole := do.SysRole{
		Id:       req.Id,
		RoleName: req.RoleName,
		RoleCode: req.RoleCode,
		RoleDesc: req.RoleDesc,
		Status:   req.Status,
	}
	_, err = roleModel.Data(sysRole).Where("id", sysRole.Id).Update()
	if err != nil {
		return nil, gerror.New("修改角色信息失败")
	}
	return
}

func (s sRole) UpdateRoleMenu(ctx context.Context, req *v1.UpdateRoleMenuReq) (res *v1.UpdateRoleMenuRes, err error) {
	//角色拥有菜单Ids
	var SysMenuIds []int64
	var SysRoleMenus []entity.SysRoleMenu
	var DeleteMenuIds []int64
	var AddMenuIds []int64
	roleMenuModel := dao.SysRoleMenu.Ctx(ctx)
	//查询角色拥有的菜单Ids
	_ = roleMenuModel.Where("role_id", req.RoleId).Scan(&SysRoleMenus)
	for _, roleMenu := range SysRoleMenus {
		SysMenuIds = append(SysMenuIds, roleMenu.MenuId)
	}
	//查询传递过来的RoleIds,里面不包含的Id(需要删除)
	DeleteMenuIds = FindMissingIds(SysMenuIds, req.MenuIds)
	AddMenuIds = FindMissingIds(req.MenuIds, SysMenuIds)
	//删除角色菜单 | 添加角色菜单
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		//删除
		_, err = roleMenuModel.Where("role_id", req.RoleId).WhereIn("menu_id", DeleteMenuIds).Delete()
		if err != nil {
			return err
		}
		for _, menuId := range AddMenuIds {
			var roleMenu *entity.SysRoleMenu
			_ = roleMenuModel.Unscoped().Where("role_id", req.RoleId).Where("menu_id", menuId).Scan(&roleMenu)
			if roleMenu != nil {
				//修改当前角色菜单删除状态为0
				roleMenu.IsDeleted = consts.ZERO_NUMBER
				_, err2 := roleMenuModel.Unscoped().Data(roleMenu).Where("id", roleMenu.Id).Where("role_id", req.RoleId).Update()
				if err2 != nil {
					return err2
				}
			} else {
				//新增角色菜单
				roleMenu := &entity.SysRoleMenu{
					RoleId: req.RoleId,
					MenuId: menuId,
				}
				_, err3 := roleMenuModel.Insert(roleMenu)
				if err3 != nil {
					return err3
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, gerror.New("用户角色菜单权限添加失败")
	}
	return
}

func (s sRole) GetPermissionTree(ctx context.Context, req *v1.GetPermissionTreeReq) (res []vo.SysPermissionTreeVo, err error) {
	var SysPermissions []entity.SysPermission
	var PermissionTreeVos []vo.SysPermissionTreeVo
	rolePermissionModel := dao.SysPermission.Ctx(ctx)
	_ = rolePermissionModel.Scan(&SysPermissions)
	PermissionTreeVos = buildTree(SysPermissions)
	return PermissionTreeVos, nil
}

func (s sRole) GetPermissionIdsByRoleId(ctx context.Context, req *v1.GetPermissionIdsByRoleIdReq) (res []int64, err error) {
	var SysRolePermissions []entity.SysRolePermission
	var permissionIds []int64
	rolePermissionModel := dao.SysRolePermission.Ctx(ctx)
	_ = rolePermissionModel.Where("role_id", req.RoleId).Scan(&SysRolePermissions)
	for _, rolePermission := range SysRolePermissions {
		permissionIds = append(permissionIds, rolePermission.PermissionId)
	}
	return permissionIds, nil
}

func (s sRole) UpdatePermissionIdsByRoleId(ctx context.Context, req *v1.UpdatePermissionIdsByRoleIdReq) (res *v1.UpdatePermissionIdsByRoleIdRes, err error) {
	//角色拥有按钮Ids
	var SysPermissionIds []int64
	var SysRolePermission []entity.SysRolePermission
	var DeletePermissionIds []int64
	var AddPermissionIds []int64
	rolePermissionModel := dao.SysRolePermission.Ctx(ctx)
	permissionModel := dao.SysPermission.Ctx(ctx)
	//查询角色拥有的菜单Ids
	_ = rolePermissionModel.Where("role_id", req.RoleId).Scan(&SysRolePermission)
	for _, rolePermission := range SysRolePermission {
		SysPermissionIds = append(SysPermissionIds, rolePermission.PermissionId)
	}
	//查询传递过来的PermissionIds,里面不包含的Id(需要删除)
	DeletePermissionIds = FindMissingIds(SysPermissionIds, req.PermissionIds)
	AddPermissionIds = FindMissingIds(req.PermissionIds, SysPermissionIds)
	//删除角色菜单 | 添加角色菜单
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		//删除
		_, err = rolePermissionModel.Where("role_id", req.RoleId).WhereIn("permission_id", DeletePermissionIds).Delete()
		if err != nil {
			return err
		}
		for _, permissionId := range AddPermissionIds {
			var rolePermission *entity.SysRolePermission
			_ = rolePermissionModel.Unscoped().Where("role_id", req.RoleId).Where("permission_id", permissionId).Scan(&rolePermission)
			if rolePermission != nil {
				//修改当前角色菜单删除状态为0
				rolePermission.IsDeleted = consts.ZERO_NUMBER
				_, err2 := rolePermissionModel.Unscoped().Data(rolePermission).Where("id", rolePermission.Id).Update()
				if err2 != nil {
					return err2
				}
			} else {
				//校验当前菜单是否存在
				count, _ := permissionModel.Where("id", permissionId).Count()
				if count <= 0 {
					continue
				}
				//新增角色菜单
				rolePermission := &entity.SysRolePermission{
					RoleId:       req.RoleId,
					PermissionId: permissionId,
				}
				_, err3 := rolePermissionModel.Insert(rolePermission)
				if err3 != nil {
					return err3
				}
			}
		}
		return nil
	})
	return
}

// buildTree 构建菜单树
func buildTree(SysPermissions []entity.SysPermission) []vo.SysPermissionTreeVo {
	//构建头节点
	var PermissionTree []vo.SysPermissionTreeVo

	for _, permission := range SysPermissions {
		//查看有没有Label一样的头节点
		isTemp := false
		for _, permissionVo := range PermissionTree {
			if permissionVo.Label == permission.MenuName {
				//相同的头节点 返回
				isTemp = true
			}
		}
		if isTemp {
			continue
		}
		//构建头节点
		treeVo := vo.SysPermissionTreeVo{
			Label:    permission.MenuName,
			Value:    int64(rand.N(1000)),
			Children: []vo.SysPermissionTreeVo{},
		}
		//构建头节点的子节点
		for _, permissionChildren := range SysPermissions {
			//构建头节点的子节点
			if treeVo.Label == permissionChildren.MenuName {
				//添加子节点
				childrenVo := vo.SysPermissionTreeVo{
					Label:    permissionChildren.Description,
					Value:    permissionChildren.Id,
					Children: []vo.SysPermissionTreeVo{},
				}
				treeVo.Children = append(treeVo.Children, childrenVo)
			}
		}
		PermissionTree = append(PermissionTree, treeVo)
	}
	return PermissionTree
}

// FindMissingIds 找出在 MasterMenuIds 中存在，但在 ServantMenuIds 中不存在的元素
func FindMissingIds(MasterMenuIds, ServantMenuIds []int64) []int64 {
	// 创建一个映射，用于快速查找 ServantMenuIds 中的元素
	idMap := make(map[int64]bool)
	for _, id := range ServantMenuIds {
		idMap[id] = true
	}

	var missingIds []int64
	// 遍历 MasterMenuIds 切片
	for _, id := range MasterMenuIds {
		// 如果在 ServantMenuIds 中不存在该元素，则添加到结果切片中
		if _, exists := idMap[id]; !exists {
			missingIds = append(missingIds, id)
		}
	}
	return missingIds
}
