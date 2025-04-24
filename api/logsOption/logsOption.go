// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package logsOption

import (
	"context"

	"Ba-Server/api/logsOption/v1"
)

type ILogsOptionV1 interface {
	GetMonOptionLogsPage(ctx context.Context, req *v1.GetMonOptionLogsPageReq) (res *v1.GetMonOptionLogsPageRes, err error)
}
