package domain

// RPage 自定义分页返回类
type RPage struct {
	// 当前页
	Current int `json:"current" yaml:"current" description:"当前页"`
	// 每页条数
	Size int `json:"size" yaml:"size" description:"每页条数"`
	// 总页数
	Pages int `json:"pages" yaml:"pages" description:"总页数"`
	// 总条数
	Total int `json:"total" yaml:"total" description:"总条数"`
	// 数据列表
	Records interface{} `json:"records" yaml:"records" description:"数据列表"`
}
