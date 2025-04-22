package logsFile

import (
	"Ba-Server/internal/model/domain"
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/logsFile/v1"
)

func (c *ControllerV1) GetMonFileLogsPage(ctx context.Context, req *v1.GetMonFileLogsPageReq) (res *v1.GetMonFileLogsPageRes, err error) {
	total, records, err := service.LogsFile().GetMonFileLogsPage(ctx, req)

	//发送错误
	if err != nil {
		return nil, gerror.New("获取角色分页异常")
	}

	//构建返回对象
	rPage := &domain.RPage{
		Current: req.Current,
		Size:    req.Size,
		Pages:   (total + req.Size - 1) / req.Size,
		Total:   total,
		Records: records,
	}
	return (*v1.GetMonFileLogsPageRes)(rPage), nil
}
