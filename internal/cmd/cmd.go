package cmd

import (
	"Ba-Server/internal/controller/auth"
	"Ba-Server/internal/controller/menu"
	"Ba-Server/internal/controller/role"
	"Ba-Server/internal/controller/route"
	"Ba-Server/internal/controller/user"
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 启动GToken
			gfToken, err := StartBackendGToken()
			if err != nil {
				return err
			}
			s.Group("/route", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, service.Middleware().ResponseHandler)

				//GToken 鉴权
				err := gfToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}
				group.Bind(
					route.NewV1(),
				)
			})
			s.Group("/auth", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, service.Middleware().ResponseHandler)
				//GToken 鉴权
				err := gfToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}
				group.Bind(
					auth.NewV1(),
				)
			})
			s.Group("/sysUser", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, service.Middleware().ResponseHandler)
				//GToken 鉴权
				err := gfToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}
				group.Bind(
					user.NewV1(),
				)
			})
			s.Group("/sysRole", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, service.Middleware().ResponseHandler)
				//GToken 鉴权
				err := gfToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}
				group.Bind(
					role.NewV1(),
				)
			})
			s.Group("/sysMenu", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse, service.Middleware().ResponseHandler)
				//GToken 鉴权
				err := gfToken.Middleware(ctx, group)
				if err != nil {
					panic(err)
				}
				group.Bind(
					menu.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
