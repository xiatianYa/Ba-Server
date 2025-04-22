package logsFile

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/logsFile/v1"
)

func (c *ControllerV1) RemoveLogsFile(ctx context.Context, req *v1.RemoveLogsFileReq) (res *v1.RemoveLogsFileRes, err error) {
	out, err := service.LogsFile().RemoveLogsFile(ctx, req)
	if err != nil {
		return nil, gerror.New("删除文件日志失败" + err.Error())
	}
	return (*v1.RemoveLogsFileRes)(&out), err
}
