package domain

// Meta 路由元数据
type Meta struct {
	// 菜单标题
	Title string `json:"title" yaml:"title" description:"菜单标题"`
	// 多语言标题
	I18nKey string `json:"i18nKey" yaml:"i18nKey" description:"多语言标题"`
	// 路由的角色列表
	Roles []string `json:"roles" yaml:"roles" description:"路由的角色列表"`
	// 缓存路由
	KeepAlive bool `json:"keepAlive" yaml:"keepAlive" description:"缓存路由"`
	// 是否为常量路由
	Constant bool `json:"constant" yaml:"constant" description:"是否为常量路由"`
	// 菜单图标
	Icon string `json:"icon" yaml:"icon" description:"菜单图标"`
	// 本地图标
	LocalIcon string `json:"localIcon" yaml:"localIcon" description:"本地图标"`
	// 排序值
	Order int `json:"order" yaml:"order" description:"排序值"`
	// 外链
	Href string `json:"href" yaml:"href" description:"外链"`
	// 是否隐藏
	HideInMenu bool `json:"hideInMenu" yaml:"hideInMenu" description:"是否隐藏"`
	// 活动菜单
	ActiveMenu string `json:"activeMenu" yaml:"activeMenu" description:"活动菜单"`
	// 支持多标签
	MultiTab bool `json:"multiTab" yaml:"multiTab" description:"支持多标签"`
	// 固定在页签中的序号
	FixedIndexInTab int `json:"fixedIndexInTab" yaml:"fixedIndexInTab" description:"固定在页签中的序号"`
	// 路由查询参数
	Query []KVPairs `json:"query" yaml:"query" description:"路由查询参数"`
	// 路由权限按钮
	Permissions []string `json:"permissions" yaml:"permissions" description:"路由权限按钮"`
}
