package user

import (
	v1 "Ba-Server/api/user/v1"
	"Ba-Server/internal/consts"
	"Ba-Server/internal/dao"
	"Ba-Server/internal/model/do"
	"Ba-Server/internal/model/entity"
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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
	pageQuery = pageQuery.OmitEmpty().WhereLike("user_name", "%"+req.UserName+"%")
	pageQuery = pageQuery.OmitEmpty().Where("user_gender", req.UserGender)
	pageQuery = pageQuery.OmitEmpty().WhereLike("nick_name", "%"+req.NickName+"%")
	pageQuery = pageQuery.OmitEmpty().WhereLike("user_phone", "%"+req.UserPhone+"%")
	pageQuery = pageQuery.OmitEmpty().WhereLike("user_email", "%"+req.UserEmail+"%")
	pageQuery = pageQuery.OmitEmpty().Where("status", req.Status)
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

func (s sUser) SaveSysUser(ctx context.Context, req *v1.SaveSysUserReq) (res *v1.SaveSysUserRes, err error) {
	//创建用户模型
	userModel := dao.SysUser.Ctx(ctx)
	//查询该用户名有没有存在
	one, _ := userModel.Where("user_name", req.UserName).Exist()
	if one != false {
		return nil, gerror.New("用户已存在,请修改用户名")
	}
	//配置用户密码
	inputPassword := req.PassWord + consts.Salt
	hash := sha256.Sum256([]byte(inputPassword))
	calculatedHash := hex.EncodeToString(hash[:])
	sysUser := do.SysUser{}
	sysUser.Password = calculatedHash
	//配置用户密码盐
	sysUser.Salt = consts.Salt
	//配置用户其他数据
	sysUser.UserName = req.UserName
	sysUser.NickName = req.NickName
	sysUser.UserEmail = req.UserEmail
	sysUser.UserGender = req.UserGender
	sysUser.UserPhone = req.UserPhone
	sysUser.Status = req.Status
	var userId int64
	userId, err = userModel.InsertAndGetId(sysUser)
	//给用户分配角色
	var sysRoles []entity.SysRole
	roleModel := dao.SysRole.Ctx(ctx)
	userRoleModel := dao.SysUserRole.Ctx(ctx)
	_ = roleModel.WhereIn("role_code", req.UserRoles).Scan(&sysRoles)
	for _, role := range sysRoles {
		sysRole := entity.SysUserRole{
			UserId: userId,
			RoleId: role.Id,
			Status: consts.ONE,
		}
		_, err = userRoleModel.Insert(&sysRole)
		if err != nil {
			return nil, gerror.New("用户权限添加失败")
		}
	}
	if err != nil {
		return nil, gerror.New("用户添加失败")
	}
	return
}

func (s sUser) RemoveSysUserByIds(ctx context.Context, req *v1.RemoveSysUserByIdsReq) (res *v1.RemoveSysUserByIdsRes, err error) {
	//删除当前用户
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		userModel := dao.SysUser.Ctx(ctx)
		userRoleModel := dao.SysUserRole.Ctx(ctx)
		_, err = userModel.WhereIn("id", req.Ids).Delete()
		if err != nil {
			return err
		}
		//删除用户下的角色对应列表
		_, err = userRoleModel.WhereIn("user_id", req.Ids).Delete()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, gerror.New("用户删除失败")
	}
	return
}
func (s sUser) RemoveSysUserById(ctx context.Context, req *v1.RemoveSysUserByIdReq) (res *v1.RemoveSysUserByIdRes, err error) {
	//删除当前用户
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		userModel := dao.SysUser.Ctx(ctx)
		userRoleModel := dao.SysUserRole.Ctx(ctx)
		_, err = userModel.Where("id", req.Id).Delete()
		if err != nil {
			return err
		}
		//删除用户下的角色对应列表
		_, err = userRoleModel.Where("user_id", req.Id).Delete()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, gerror.New("用户删除失败")
	}
	return
}

func (s sUser) UpdateSysUser(ctx context.Context, req *v1.UpdateSysUserReq) (res *v1.UpdateSysUserRes, err error) {
	sysUser := do.SysUser{}
	//配置用户其他数据
	sysUser.Id = req.Id
	sysUser.UserName = req.UserName
	sysUser.NickName = req.NickName
	sysUser.UserEmail = req.UserEmail
	sysUser.UserGender = req.UserGender
	sysUser.UserPhone = req.UserPhone
	sysUser.Status = req.Status
	//修改用户信息
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		//创建用户模型
		userModel := dao.SysUser.Ctx(ctx)
		roleModel := dao.SysRole.Ctx(ctx)
		userRoleModel := dao.SysUserRole.Ctx(ctx)
		_, err = userModel.Data(sysUser).Where("id", sysUser.Id).Update()
		if err != nil {
			return err
		}
		//给用户分配角色
		var sysRoles []entity.SysRole
		_ = roleModel.WhereIn("role_code", req.UserRoles).Scan(&sysRoles)
		for _, role := range sysRoles {
			sysRole := entity.SysUserRole{
				UserId: req.Id,
				RoleId: role.Id,
				Status: consts.ONE,
			}
			_, err = userRoleModel.Insert(&sysRole)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, gerror.New("用户信息修改失败")
	}
	return
}
