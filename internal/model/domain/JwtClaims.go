package domain

import "github.com/golang-jwt/jwt/v5"

// JwtClaims 定义一个自定义的Claims结构体，用于存储JWT中的数据
type JwtClaims struct {
	// 用户id
	ID int64 `json:"id"`
	// 账号
	UserName string `json:"userName"`
	// 昵称
	NickName string `json:"nickName"`
	// 头像
	Avatar string `json:"avatar"`
	// 邮箱
	UserEmail string `json:"userEmail"`
	// 手机
	UserPhone string `json:"userPhone"`
	// 性别 0保密 1男 2女
	UserGender int `json:"userGender"`
	// 角色IDs
	RoleIds []int64 `json:"roleIds"`
	// 角色Codes
	RoleCodes []string `json:"roleCodes"`
	// 支持
	jwt.RegisteredClaims
}
