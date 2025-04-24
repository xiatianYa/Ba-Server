package logsOption

import (
	"Ba-Server/internal/model/domain"
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/logsOption/v1"
)

func (c *ControllerV1) GetMonOptionLogsPage(ctx context.Context, req *v1.GetMonOptionLogsPageReq) (res *v1.GetMonOptionLogsPageRes, err error) {
	total, records, err := service.LogsOption().GetMonLogsOptionPage(ctx, req)

	//发送错误
	if err != nil {
		return nil, gerror.New("获取操作日志分页异常" + err.Error())
	}

	//构建返回对象
	rPage := &domain.RPage{
		Current: req.Current,
		Size:    req.Size,
		Pages:   (total + req.Size - 1) / req.Size,
		Total:   total,
		Records: records,
	}
	return (*v1.GetMonOptionLogsPageRes)(rPage), nil
}
