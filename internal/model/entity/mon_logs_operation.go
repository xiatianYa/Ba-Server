// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MonLogsOperation is the golang structure for table mon_logs_operation.
type MonLogsOperation struct {
	Id            int64       `json:"id"            orm:"id"             ` // 主键ID
	Ip            string      `json:"ip"            orm:"ip"             ` // IP
	IpAddr        string      `json:"ipAddr"        orm:"ip_addr"        ` // IP所属地
	UserId        int64       `json:"userId"        orm:"user_id"        ` // 此操作用户ID
	UserAgent     string      `json:"userAgent"     orm:"user_agent"     ` // 登录代理
	RequestUri    string      `json:"requestUri"    orm:"request_uri"    ` // 请求URI
	RequestPath   string      `json:"requestPath"   orm:"request_path"   ` // 请求PATH
	RequestMethod string      `json:"requestMethod" orm:"request_method" ` // 请求方式
	MethodParams  string      `json:"methodParams"  orm:"method_params"  ` // 请求参数
	UseTime       string      `json:"useTime"       orm:"use_time"       ` // 请求耗时
	CreateTime    *gtime.Time `json:"createTime"    orm:"create_time"    ` // 创建时间
	UpdateTime    *gtime.Time `json:"updateTime"    orm:"update_time"    ` // 修改时间
	IsDeleted     uint        `json:"isDeleted"     orm:"is_deleted"     ` // 是否删除(0:否,1:是)
}
