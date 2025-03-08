package auth

import (
	v1 "Ba-Server/api/auth/v1"
	"Ba-Server/internal/consts"
	"Ba-Server/internal/model/domain"
	"Ba-Server/internal/model/entity"
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type (
	sAuth struct{}
)

func init() {
	service.RegisterAuth(New())
}

func New() service.IAuth {
	return &sAuth{}
}

func (s sAuth) UserNameLogin(ctx context.Context, req *v1.LoginReq) (constantRoutes map[string]interface{}, err error) {
	//查询当前用户
	userModel := g.Model("sys_user")
	var sysUser entity.SysUser
	one, err := userModel.Where("user_name=?", req.Username).One()
	if err != nil || one == nil {
		panic(fmt.Sprintf("查找不到用户名 %s", req.Username))
	}
	gconv.Scan(one, &sysUser)

	//校验用户是否被禁止登录
	if string(consts.ZERO) == sysUser.Status {
		panic(fmt.Sprintf("当前用户 %s 已被禁止登录", req.Username))
	}

	// 密码拼接
	inputPassword := req.Password + sysUser.Salt

	// 计算 SHA-256 哈希值
	hash := sha256.Sum256([]byte(inputPassword))
	calculatedHash := hex.EncodeToString(hash[:])

	// 密码比对
	if calculatedHash != sysUser.Password {
		panic(fmt.Sprintf("当前用户 %s 密码错误", req.Username))
	}

	// 更新用户登录时间
	_, err = userModel.Where("id", sysUser.Id).Data(map[string]interface{}{
		"last_login_time": gtime.Now(),
	}).Update()

	if err != nil {
		panic("用户登录时间更新失败")
	}

	//生成Token 返回
	token, err := saveUserToSession(sysUser)
	if err != nil {
		panic("token生成失败")
	}
	data := map[string]interface{}{
		"token":        token,
		"refreshToken": token,
	}
	return data, nil
}

func (s sAuth) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (constantRoutes vo.SysUserInfoVo, err error) {
	return
}

// 根据用户生成Token
func saveUserToSession(sysUser entity.SysUser) (resultToken string, err error) {
	// 用户用户的角色对应列表
	var sysUserRoles []entity.SysUserRole
	// 用户拥有的角色列表
	var sysRoles []entity.SysRole
	// 创建一个新的切片来存储 roleId
	var roleIds []int64
	// 创建一个新的切片来存储 roleCode
	var roleCodes []string
	//查询用户拥有的RoleIds 和 RoleCodes
	userRole := g.Model("sys_role")
	userRoleModel := g.Model("sys_user_role")
	//获取用户角色列表
	_ = userRoleModel.Where("user_id=?", sysUser.Id).Scan(&sysUserRoles)
	// 遍历 sysUserRoles 切片，将每个 Role 结构体的 roleId 字段添加到 roleIds 切片中
	for _, role := range sysUserRoles {
		roleIds = append(roleIds, role.RoleId)
	}
	_ = userRole.Where("status", 1).WhereIn("id", roleIds).Scan(&sysRoles)
	// 遍历 sysRoles 切片，将每个 Role 结构体的 roleId 字段添加到 roleIds 切片中
	for _, role := range sysRoles {
		roleCodes = append(roleCodes, role.RoleCode)
	}
	// 生成 token
	uc := &domain.JwtClaims{
		ID:         sysUser.Id,
		UserName:   sysUser.UserName,
		NickName:   sysUser.NickName,
		Avatar:     sysUser.Avatar,
		UserEmail:  sysUser.UserEmail,
		UserPhone:  sysUser.UserPhone,
		UserGender: sysUser.UserGender,
		RoleIds:    roleIds,
		RoleCodes:  roleCodes,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)

	key := []byte(consts.JwtKey)

	// 使用转换后的密钥生成签名后的令牌
	return token.SignedString(key)
}
