// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MonLogsFile is the golang structure for table mon_logs_file.
type MonLogsFile struct {
	Id           int64       `json:"id"           orm:"id"             ` // 主键ID
	UserId       int64       `json:"userId"       orm:"user_id"        ` // 用户ID
	FilePath     string      `json:"filePath"     orm:"file_path"      ` // 文件物理路径
	FileUrl      string      `json:"fileUrl"      orm:"file_url"       ` // 文件网络路径
	FileType     int         `json:"fileType"     orm:"file_type"      ` // 文件类型
	FileSize     string      `json:"fileSize"     orm:"file_size"      ` // 文件大小
	ErrorMsg     string      `json:"errorMsg"     orm:"error_msg"      ` // 异常日志
	Status       string      `json:"status"       orm:"status"         ` // 上传状态(0:失败 1:成功)
	CreateUserId int64       `json:"createUserId" orm:"create_user_id" ` // 创建用户ID
	CreateTime   *gtime.Time `json:"createTime"   orm:"create_time"    ` // 创建时间
	UpdateUserId int64       `json:"updateUserId" orm:"update_user_id" ` // 修改用户ID
	UpdateTime   *gtime.Time `json:"updateTime"   orm:"update_time"    ` // 修改时间
	IsDeleted    uint        `json:"isDeleted"    orm:"is_deleted"     ` // 是否删除(0:否,1:是)
}
