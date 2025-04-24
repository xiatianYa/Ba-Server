package vo

import "github.com/gogf/gf/v2/os/gtime"

// MonLogsOptionVo is the golang structure for table mon_logs_file.
type MonLogsOptionVo struct {
	Id            int64       `json:"id"`            // 主键ID
	Ip            string      `json:"ip"`            // IP
	UserId        int64       `json:"userId"`        //用户Id
	UserName      string      `json:"userName"`      //用户名称
	IpAddr        string      `json:"ipAddr"`        // IP所属地
	UserAgent     string      `json:"userAgent"`     // 登录代理
	RequestUri    string      `json:"requestUri"`    // 请求URI
	RequestPath   string      `json:"requestPath"`   // 请求PATH
	RequestMethod string      `json:"requestMethod"` // 请求方式
	MethodParams  string      `json:"methodParams"`  // 请求参数
	UseTime       string      `json:"useTime"`       // 请求耗时
	CreateTime    *gtime.Time `json:"createTime"`    // 创建时间
}
