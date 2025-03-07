package vo

import "Ba-Server/internal/model/domain"

// SysUserRouteVO 用户权限路由对象
type SysUserRouteVO struct {
	// 首页
	Home string `json:"home" yaml:"home" description:"首页"`
	// 路由列表
	Routes []domain.Route `json:"routes" yaml:"routes" description:"路由列表"`
}
