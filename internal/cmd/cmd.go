package cmd

import (
	"Ba-Server/internal/controller/auth"
	"Ba-Server/internal/controller/dict"
	"Ba-Server/internal/controller/file"
	"Ba-Server/internal/controller/logsFile"
	"Ba-Server/internal/controller/logsOption"
	"Ba-Server/internal/controller/menu"
	"Ba-Server/internal/controller/role"
	"Ba-Server/internal/controller/route"
	"Ba-Server/internal/controller/user"
	"Ba-Server/internal/logic/middleware"
	"Ba-Server/internal/service"
	fileUtil "Ba-Server/utility/file"
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
			// createDirIfNotExists
			_ = fileUtil.CreateDirIfNotExists(fileUtil.FilePath)
			s.SetServerRoot(fileUtil.FilePath)
			s.AddStaticPath("/statics", fileUtil.FilePath)
			// 启动GToken
			gfToken, _ := StartBackendGToken()
			// 全局日志拦截
			s.Use(middleware.OptionsLogMiddleWare)
			s.Group("/", func(group *ghttp.RouterGroup) {
				// 返回拦截
				group.Middleware(ghttp.MiddlewareHandlerResponse, service.Middleware().ResponseHandler)
				//GToken 鉴权
				_ = gfToken.Middleware(ctx, group)
				group.Group("/route", func(group *ghttp.RouterGroup) {
					group.Bind(
						route.NewV1(),
					)
				})
				group.Group("/auth", func(group *ghttp.RouterGroup) {
					group.Bind(
						auth.NewV1(),
					)
				})
				group.Group("/sysUser", func(group *ghttp.RouterGroup) {
					group.Bind(
						user.NewV1(),
					)
				})
				group.Group("/sysRole", func(group *ghttp.RouterGroup) {
					group.Bind(
						role.NewV1(),
					)
				})
				group.Group("/sysMenu", func(group *ghttp.RouterGroup) {
					group.Bind(
						menu.NewV1(),
					)
				})
				group.Group("/sysDict", func(group *ghttp.RouterGroup) {
					group.Bind(
						dict.NewV1(),
					)
				})
				group.Group("/sysFile", func(group *ghttp.RouterGroup) {
					group.Bind(
						file.NewV1(),
					)
				})
				group.Group("/monLogsFile", func(group *ghttp.RouterGroup) {
					group.Bind(
						logsFile.NewV1(),
					)
				})
				group.Group("/monLogsOperation", func(group *ghttp.RouterGroup) {
					group.Bind(
						logsOption.NewV1(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
