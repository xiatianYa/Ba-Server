// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	v1 "Ba-Server/api/logsOption/v1"
	"Ba-Server/internal/model/vo"
	"context"
)

type (
	ILogsOption interface {
		GetMonLogsOptionPage(ctx context.Context, req *v1.GetMonOptionLogsPageReq) (total int, records []vo.MonLogsOptionVo, err error)
	}
)

var (
	localLogsOption ILogsOption
)

func LogsOption() ILogsOption {
	if localLogsOption == nil {
		panic("implement not found for interface ILogsOption, forgot register?")
	}
	return localLogsOption
}

func RegisterLogsOption(i ILogsOption) {
	localLogsOption = i
}
