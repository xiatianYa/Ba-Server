package cmd

import (
	"Ba-Server/internal/consts"
	"Ba-Server/internal/dao"
	"Ba-Server/internal/model/entity"
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
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

	userInfoVo, err := service.User().GetUserInfoVo(ctx, sysUser)
	if err != nil {
		response.ErrorExit(r, "获取用户权限信息失败")
	}

	return consts.GTokenPrefix + strconv.FormatInt(sysUser.Id, 10), userInfoVo
}

// 登录之后的函数
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	ctx := context.TODO()
	if !respData.Success() {
		respData.Code = consts.ERROR
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = consts.SUCCESS
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
	if containsSubstring(sysUserInfoVo.Buttons, r.URL.Path) {
		response.AuthPermission(r)
		return
	}
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

// 校验切片是否包含某个字符
func containsSubstring(arr []string, target string) bool {
	for _, item := range arr {
		if strings.HasPrefix(target, item) {
			return false
		}
	}
	return true
}
