package vo

import "github.com/gogf/gf/v2/os/gtime"

// MonLogsFileVo is the golang structure for table mon_logs_file.
type MonLogsFileVo struct {
	Id         int64       `json:"id"`         // 主键ID
	UserId     int64       `json:"userId"`     // 用户ID
	UserName   string      `json:"userName"`   // 用户名称
	FilePath   string      `json:"filePath"`   // 文件物理路径
	FileUrl    string      `json:"fileUrl"`    // 文件网络路径
	FileType   int         `json:"fileType"`   // 文件类型
	FileSize   int         `json:"fileSize"`   // 文件大小
	ErrorMsg   string      `json:"errorMsg"`   // 异常日志
	Status     string      `json:"status"`     // 上传状态(0:失败 1:成功)
	CreateTime *gtime.Time `json:"createTime"` // 创建时间
}
