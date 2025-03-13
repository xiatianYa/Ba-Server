package role

import (
	v1 "Ba-Server/api/role/v1"
	"Ba-Server/internal/consts"
	"Ba-Server/internal/dao"
	"Ba-Server/internal/model/do"
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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
