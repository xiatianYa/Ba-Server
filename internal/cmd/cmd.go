package cmd

import (
	"Ba-Server/internal/consts"
	"Ba-Server/internal/model/entity"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"strconv"

	"Ba-Server/internal/controller/auth"
	"Ba-Server/internal/controller/hello"
	"Ba-Server/internal/controller/route"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 启动GToken
			gfToken := &gtoken.GfToken{
				LoginPath:        "/login",
				LoginBeforeFunc:  LoginFunc,
				LogoutPath:       "/user/logout",
				AuthExcludePaths: g.SliceStr{},
				MultiLogin:       true,
			}
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
				)
			})
			s.Group("/route", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					route.NewV1(),
				)
			})
			s.Group("/auth", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				//GToken
				err := gfToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}
				group.Bind(
					auth.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)

// TODO
func LoginFunc(r *ghttp.Request) (string, interface{}) {
	Username := r.Get("userName").String()
	Password := r.Get("password").String()
	//查询当前用户
	userModel := g.Model("sys_user")
	var sysUser entity.SysUser
	one, err := userModel.Where("user_name=?", Username).One()
	if err != nil || one == nil {
		panic(fmt.Sprintf("查找不到用户名 %s", Username))
	}
	gconv.Scan(one, &sysUser)

	//校验用户是否被禁止登录
	if consts.ZERO == sysUser.Status {
		panic(fmt.Sprintf("当前用户 %s 已被禁止登录", Username))
	}

	// 密码拼接
	inputPassword := Password + sysUser.Salt

	// 计算 SHA-256 哈希值
	hash := sha256.Sum256([]byte(inputPassword))
	calculatedHash := hex.EncodeToString(hash[:])

	// 密码比对
	if calculatedHash != sysUser.Password {
		panic(fmt.Sprintf("当前用户 %s 密码错误", Username))
	}

	// 更新用户登录时间
	_, err = userModel.Where("id", sysUser.Id).Data(map[string]interface{}{
		"last_login_time": gtime.Now(),
	}).Update()

	if err != nil {
		panic("用户登录时间更新失败")
	}

	return consts.GTokenPrefix + strconv.FormatInt(sysUser.Id, 10), sysUser
}
