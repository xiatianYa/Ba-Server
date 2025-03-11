package user

import (
	v1 "Ba-Server/api/user/v1"
	"Ba-Server/internal/dao"
	"Ba-Server/internal/model/entity"
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

func (s sUser) GetSysUserPage(ctx context.Context, req *v1.GetSysUserPageReq) (total int, records []*vo.SysUserVo, err error) {
	// 创建指针切片
	var result []*vo.SysUserVo

	userModel := dao.SysUser.Ctx(ctx)
	userRoleModel := dao.SysUserRole.Ctx(ctx)
	roleModel := dao.SysRole.Ctx(ctx)
	pageQuery := userModel.Page(req.Current, req.Size)

	// 处理各个查询条件，只添加非空的条件
	if req.UserName != "" {
		pageQuery = pageQuery.Where("user_name like ?", "%"+req.UserName+"%")
	}
	if req.UserGender != 0 {
		pageQuery = pageQuery.Where("user_gender = ?", req.UserGender)
	}
	if req.NickName != "" {
		pageQuery = pageQuery.Where("nick_name like ?", "%"+req.NickName+"%")
	}
	if req.UserPhone != "" {
		pageQuery = pageQuery.Where("user_phone like ?", "%"+req.UserPhone+"%")
	}
	if req.UserEmail != "" {
		pageQuery = pageQuery.Where("user_email like ?", "%"+req.UserEmail+"%")
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
		records = []*vo.SysUserVo{}
		return
	}

	//获取用户角色标识列表
	for _, sysUser := range result {
		// 创建一个新的切片来存储 roleIds
		var roleIds []int64
		// 用户拥有的角色列表
		var sysUserRoles []entity.SysUserRole
		var sysRoles []entity.SysRole
		// 创建一个新的切片来存储 roleCode
		var roleCodes []string
		_ = userRoleModel.Where("user_id", sysUser.Id).Scan(&sysUserRoles)
		// 用户拥有的角色标识
		for _, sysUserRole := range sysUserRoles {
			roleIds = append(roleIds, sysUserRole.RoleId)
		}
		//获取用户角色标识列表
		_ = roleModel.WhereIn("id", roleIds).Scan(&sysRoles)
		for _, sysRole := range sysRoles {
			roleCodes = append(roleCodes, sysRole.RoleCode)
		}
		sysUser.UserRoles = roleCodes
	}
	records = result
	return
}

func (s sUser) GetUserInfoVo(ctx context.Context, sysUser entity.SysUser) (sysUserInfoVo *vo.SysUserInfoVo, err error) {
	// 用户用户的角色对应列表
	var sysUserRoles []entity.SysUserRole
	// 用户拥有的角色列表
	var sysRoles []entity.SysRole
	// 用户拥有的角色权限列表
	var SysRolePermissions []entity.SysRolePermission
	// 用户用户的权限列表
	var SysPermissions []entity.SysPermission
	// 创建一个新的切片来存储 roleIds
	var roleIds []int64
	// 创建一个新的切片来存储 roleCode
	var roleCodes []string
	// 创建一个新的切片来存储 roleButton
	var roleButtons []string
	// 创建一个新的切片来存储 permissionIds
	var permissionIds []int64
	//查询用户拥有的RoleIds 和 RoleCodes
	RoleModel := dao.SysRole.Ctx(ctx)
	UserRoleModel := dao.SysUserRole.Ctx(ctx)
	RolePermissionModel := dao.SysRolePermission.Ctx(ctx)
	PermissionModel := dao.SysPermission.Ctx(ctx)
	//获取用户角色列表
	_ = UserRoleModel.Where("user_id=?", sysUser.Id).Scan(&sysUserRoles)
	// 遍历 sysUserRoles 切片，将每个 Role 结构体的 roleId 字段添加到 roleIds 切片中
	for _, role := range sysUserRoles {
		roleIds = append(roleIds, role.RoleId)
	}
	_ = RoleModel.WhereIn("id", roleIds).Scan(&sysRoles)
	// 遍历 sysRoles 切片，将每个 Role 结构体的 roleCode 字段添加到 roleCodes 切片中
	for _, role := range sysRoles {
		roleCodes = append(roleCodes, role.RoleCode)
	}
	//获取用户按钮列表
	_ = RolePermissionModel.WhereIn("role_id", roleIds).Scan(&SysRolePermissions)
	// 遍历 SysRolePermissions 切片，将每个 rolePermission 结构体的 permissionId 字段添加到 permissionIds 切片中
	for _, rolePermission := range SysRolePermissions {
		permissionIds = append(permissionIds, rolePermission.PermissionId)
	}
	_ = PermissionModel.WhereIn("id", permissionIds).Scan(&SysPermissions)
	// 遍历 SysPermissions 切片，将每个 permission 结构体的 code 字段添加到 roleButtons 切片中
	for _, permission := range SysPermissions {
		roleButtons = append(roleButtons, permission.Code)
	}
	return &vo.SysUserInfoVo{
		UserID:     sysUser.Id,
		UserName:   sysUser.UserName,
		Avatar:     sysUser.Avatar,
		UserPhone:  sysUser.UserPhone,
		UserEmail:  sysUser.UserEmail,
		UserGender: sysUser.UserGender,
		RoleCodes:  roleCodes,
		Buttons:    roleButtons,
	}, nil
}
