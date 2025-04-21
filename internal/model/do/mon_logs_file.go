// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MonLogsFile is the golang structure of table mon_logs_file for DAO operations like Where/Data.
type MonLogsFile struct {
	g.Meta       `orm:"table:mon_logs_file, do:true"`
	Id           interface{} // 主键ID
	UserId       interface{} // 用户ID
	UserName     interface{} // 用户名称
	FilePath     interface{} // 文件物理路径
	FileUrl      interface{} // 文件网络路径
	FileType     interface{} // 文件类型
	FileSize     interface{} // 文件大小
	ErrorMsg     interface{} // 异常日志
	Status       interface{} // 上传状态(0:失败 1:成功)
	CreateUserId interface{} // 创建用户ID
	CreateTime   *gtime.Time // 创建时间
	UpdateUserId interface{} // 修改用户ID
	UpdateTime   *gtime.Time // 修改时间
	IsDeleted    interface{} // 是否删除(0:否,1:是)
}
