// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MonLogsOperation is the golang structure of table mon_logs_operation for DAO operations like Where/Data.
type MonLogsOperation struct {
	g.Meta        `orm:"table:mon_logs_operation, do:true"`
	Id            interface{} // 主键ID
	Ip            interface{} // IP
	IpAddr        interface{} // IP所属地
	UserAgent     interface{} // 登录代理
	RequestUri    interface{} // 请求URI
	RequestPath   interface{} // 请求PATH
	RequestMethod interface{} // 请求方式
	MethodParams  interface{} // 请求参数
	UseTime       interface{} // 请求耗时
	CreateTime    *gtime.Time // 创建时间
	UpdateTime    *gtime.Time // 修改时间
	IsDeleted     interface{} // 是否删除(0:否,1:是)
}
