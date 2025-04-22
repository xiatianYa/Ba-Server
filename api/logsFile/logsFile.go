// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package logsFile

import (
	"context"

	"Ba-Server/api/logsFile/v1"
)

type ILogsFileV1 interface {
	GetMonFileLogsPage(ctx context.Context, req *v1.GetMonFileLogsPageReq) (res *v1.GetMonFileLogsPageRes, err error)
}
