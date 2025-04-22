package logsFile

import (
	v1 "Ba-Server/api/logsFile/v1"
	"Ba-Server/internal/dao"
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"
)

type (
	sLogsFile struct{}
)

func init() {
	service.RegisterLogsFile(New())
}

func New() service.ILogsFile {
	return &sLogsFile{}
}

func (s sLogsFile) GetMonFileLogsPage(ctx context.Context, req *v1.GetMonFileLogsPageReq) (total int, records []vo.MonLogsFileVo, err error) {
	// 创建指针切片
	var result []vo.MonLogsFileVo

	monLogsFileModel := dao.MonLogsFile.Ctx(ctx)
	pageQuery := monLogsFileModel.Page(req.Current, req.Size)
	pageQuery = monLogsFileModel.OmitEmpty().Where("user_id", req.UserId)
	if req.StartTime != "" || req.EndTime != "" {
		pageQuery = monLogsFileModel.OmitEmpty().WhereBetween("create_time", req.StartTime, req.EndTime)
	}
	err = pageQuery.ScanAndCount(&result, &total, true)

	if err != nil {
		return 0, nil, err
	}

	// 如果列表为空
	if result == nil {
		records = []vo.MonLogsFileVo{}
		return
	}
	records = result

	return
}
