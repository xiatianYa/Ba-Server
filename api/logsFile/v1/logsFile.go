package v1

import (
	"Ba-Server/internal/model/domain"
	"github.com/gogf/gf/v2/frame/g"
)

type GetMonFileLogsPageReq struct {
	g.Meta    `path:"/page" tags:"page" method:"get" summary:"获取文件日志分页"`
	Current   int    `v:"required" json:"current"`
	Size      int    `v:"required" json:"size"`
	UserId    int64  `json:"userName"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type GetMonFileLogsPageRes domain.RPage

type RemoveLogsFileReq struct {
	g.Meta `path:"/remove/{id}" tags:"remove" method:"delete" summary:"删除单个文件日志"`
	Id     int64 `v:"required" dc:"logs file id"`
}

type RemoveLogsFileRes bool
