// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "Ba-Server/api/logsFile/v1"
	"Ba-Server/internal/model/vo"
	"context"
)

type (
	ILogsFile interface {
		GetMonFileLogsPage(ctx context.Context, req *v1.GetMonFileLogsPageReq) (total int, records []vo.MonLogsFileVo, err error)
		// RemoveLogsFile 删除单个文件日志
		RemoveLogsFile(ctx context.Context, req *v1.RemoveLogsFileReq) (bool, error)
	}
)

var (
	localLogsFile ILogsFile
)

func LogsFile() ILogsFile {
	if localLogsFile == nil {
		panic("implement not found for interface ILogsFile, forgot register?")
	}
	return localLogsFile
}

func RegisterLogsFile(i ILogsFile) {
	localLogsFile = i
}
