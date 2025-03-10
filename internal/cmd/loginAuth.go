package cmd

import (
	"Ba-Server/internal/consts"
	"Ba-Server/internal/dao"
	"Ba-Server/internal/model/entity"
	"Ba-Server/internal/model/vo"
	"Ba-Server/utility/response"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"strconv"
	"strings"
)

// StartBackendGToken 开启GFToken
func StartBackendGToken() (gfAdminToken *gtoken.GfToken, err error) {
	// 启动GToken
	gfToken := &gtoken.GfToken{
		ServerName:       "Ba-Server",
		CacheMode:        consts.GTokenRedisCache,
		LoginPath:        "/login",
		LoginBeforeFunc:  loginFunc,
		LoginAfterFunc:   loginAfterFunc,
		LogoutPath:       "/user/logout",
		AuthExcludePaths: g.SliceStr{"/route/getConstantRoutes"},
		MultiLogin:       consts.GTokenNoMultiLogin,
		AuthAfterFunc:    authAfterFunc,
	}
	err = gfToken.Start()
	return gfToken, nil
}

// 登录校验函数
func loginFunc(r *ghttp.Request) (string, interface{}) {
	Username := r.Get("userName").String()
	Password := r.Get("password").String()
	ctx := context.TODO()
	//查询当前用户
	userModel := dao.SysUser.Ctx(ctx)
	var sysUser entity.SysUser
	one, err := userModel.Where("user_name=?", Username).One()
	if err != nil || one == nil {
		response.ErrorExit(r, fmt.Sprintf("查找不到用户名 %s", Username))
	}
	_ = gconv.Scan(one, &sysUser)

	//校验用户是否被禁止登录
	if consts.ZERO == sysUser.Status {
		response.AuthBlack(r)
	}

	// 密码拼接
	inputPassword := Password + sysUser.Salt

	// 计算 SHA-256 哈希值
	hash := sha256.Sum256([]byte(inputPassword))
	calculatedHash := hex.EncodeToString(hash[:])

	// 密码比对
	if calculatedHash != sysUser.Password {
		response.ErrorExit(r, fmt.Sprintf("当前用户 %s 密码错误", Username))
	}

	// 更新用户登陆时间
	_, err = userModel.Where("id", sysUser.Id).Data(map[string]interface{}{
		"last_login_time": gtime.Now(),
	}).Update()

	if err != nil {
		response.ErrorExit(r, "更新用户登陆时间失败")
	}

	infoVo, err := getUserInfoVo(sysUser)
	if err != nil {
		response.ErrorExit(r, "获取用户权限信息失败")
	}

	return consts.GTokenPrefix + strconv.FormatInt(sysUser.Id, 10), infoVo
}

// 登录之后的函数
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	ctx := context.TODO()
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		// 获得登录用户id
		userKey := respData.GetString("userKey")
		userId := gstr.StrEx(userKey, consts.GTokenPrefix)
		// 根据id获得登录用户其他信息
		adminInfo := entity.SysUser{}
		err := dao.SysUser.Ctx(ctx).WherePri(userId).Scan(&adminInfo)
		// 确认角色不为nil
		if err != nil {
			return
		}
		data := map[string]interface{}{
			"token":        respData.GetString("token"),
			"refreshToken": respData.GetString("token"),
		}
		response.JsonExit(r, consts.SUCCESS, "success", data)
	}
	return
}

// 权限校验函数
func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	var sysUserInfoVo vo.SysUserInfoVo
	err := gconv.Struct(respData.GetString("data"), &sysUserInfoVo)
	// 账号不存在
	if err != nil {
		response.Auth(r)
		return
	}
	//todo 根据请求路径校验用户是否有权限操作接口
	//if containsSubstring(sysUserInfoVo.Buttons, r.URL.Path) {
	//	response.AuthPermission(r)
	//	return
	//}
	//todo 这里可以写账号前置校验、是否被拉黑、有无权限等逻辑
	r.SetCtxVar(consts.UserId, sysUserInfoVo.UserID)
	r.SetCtxVar(consts.UserName, sysUserInfoVo.UserName)
	r.SetCtxVar(consts.UserAvatar, sysUserInfoVo.Avatar)
	r.SetCtxVar(consts.UserEmail, sysUserInfoVo.UserEmail)
	r.SetCtxVar(consts.UserGender, sysUserInfoVo.UserGender)
	r.SetCtxVar(consts.UserPhone, sysUserInfoVo.UserPhone)
	r.SetCtxVar(consts.RoleCodes, sysUserInfoVo.RoleCodes)
	r.SetCtxVar(consts.Buttons, sysUserInfoVo.Buttons)
	r.Middleware.Next()
}

// 获取用户权限列表函数
func getUserInfoVo(sysUser entity.SysUser) (vo.SysUserInfoVo, error) {
	ctx := context.TODO()
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
	userInfoVo := vo.SysUserInfoVo{
		UserID:     sysUser.Id,
		UserName:   sysUser.UserName,
		Avatar:     sysUser.Avatar,
		UserPhone:  sysUser.UserPhone,
		UserEmail:  sysUser.UserEmail,
		UserGender: sysUser.UserGender,
		RoleCodes:  roleCodes,
		Buttons:    roleButtons,
	}
	return userInfoVo, nil
}

// 校验切片是否包含某个字符
func containsSubstring(arr []string, target string) bool {
	for _, item := range arr {
		if strings.Contains(item, target) {
			return false
		}
	}
	return true
}
