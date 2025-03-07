package domain

// Route 路由对象
type Route struct {
	// 菜单名称(唯一)
	Name string `json:"name" yaml:"name" description:"菜单名称(唯一)"`
	// 菜单路径
	Path string `json:"path" yaml:"path" description:"菜单路径"`
	// 路由组件；详情可查阅：https://github.com/soybeanjs/elegant-router/blob/main/README.zh_CN.md
	Component string `json:"component" yaml:"component" description:"路由组件；详情可查阅：https://github.com/soybeanjs/elegant-router/blob/main/README.zh_CN.md"`
	// 路由元数据
	Meta Meta `json:"meta" yaml:"meta" description:"路由元数据"`
	// 渲染道具
	Props bool `json:"props" yaml:"props" description:"渲染道具"`
	// 子菜单
	Children []Route `json:"children,omitempty" yaml:"children,omitempty" description:"子菜单"`
}
