package route

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"Ba-Server/api/route/v1"
)

func (c *ControllerV1) GetConstantRoutes(ctx context.Context, req *v1.GetConstantRoutesReq) (res *v1.GetConstantRoutesRes, err error) {
	// 构建返回的 map
	result := make(map[string]interface{})
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
	result["data"] = data
	result["code"] = "200"
	result["msg"] = "success"
	g.RequestFromCtx(ctx).Response.WriteJson(result)
	return
}
