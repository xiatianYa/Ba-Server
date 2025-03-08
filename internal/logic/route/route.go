package route

import (
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	sRoute struct{}
)

func init() {
	service.RegisterRoute(New())
}

func New() service.IRoute {
	return &sRoute{}
}

func (s sRoute) GetConstantRoutes(ctx context.Context) (constantRoutes []map[string]interface{}, err error) {
	data := []map[string]interface{}{
		{
			"name":      "403",
			"path":      "/403",
			"component": "layout.blank$view.403",
			"props":     false,
			"meta": map[string]interface{}{
				"title":           "403",
				"i18nKey":         "route.403",
				"localIcon":       "",
				"icon":            "",
				"order":           0,
				"hideInMenu":      true,
				"activeMenu":      "",
				"multiTab":        false,
				"fixedIndexInTab": nil,
				"keepAlive":       false,
				"constant":        true,
				"href":            "",
				"query":           nil,
			},
		},
		{
			"name":      "404",
			"path":      "/404",
			"component": "layout.blank$view.404",
			"props":     false,
			"meta": map[string]interface{}{
				"title":           "404",
				"i18nKey":         "route.404",
				"localIcon":       "",
				"icon":            "",
				"order":           0,
				"hideInMenu":      true,
				"activeMenu":      "",
				"multiTab":        false,
				"fixedIndexInTab": nil,
				"keepAlive":       false,
				"constant":        true,
				"href":            "",
				"query":           nil,
			},
		},
		{
			"name":      "500",
			"path":      "/500",
			"component": "layout.blank$view.500",
			"props":     false,
			"meta": map[string]interface{}{
				"title":           "500",
				"i18nKey":         "route.500",
				"localIcon":       "",
				"icon":            "",
				"order":           0,
				"hideInMenu":      true,
				"activeMenu":      "",
				"multiTab":        false,
				"fixedIndexInTab": nil,
				"keepAlive":       false,
				"constant":        true,
				"href":            "",
				"query":           nil,
			},
		},
		{
			"name":      "login",
			"path":      "/login/:module(pwd-login|code-login|register|reset-pwd|bind-wechat)?",
			"component": "layout.blank$view.login",
			"props":     true,
			"meta": map[string]interface{}{
				"title":           "login",
				"i18nKey":         "route.login",
				"localIcon":       "",
				"icon":            "",
				"order":           0,
				"hideInMenu":      true,
				"activeMenu":      "",
				"multiTab":        false,
				"fixedIndexInTab": nil,
				"keepAlive":       false,
				"constant":        true,
				"href":            "",
				"query":           nil,
			},
		},
	}
	return data, nil
}

func (s sRoute) GetUserRoutes(ctx context.Context) (userRouteVo vo.SysUserRouteVO, err error) {
	//创建用户权限模型
	userRoleModel := g.Model("sys_user_role")
	//查询用户角色
	userRoleModel.Where("user_id=?", 1).Array()
	return
}
