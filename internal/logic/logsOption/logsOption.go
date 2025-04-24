package logsOption

import (
	v1 "Ba-Server/api/logsOption/v1"
	"Ba-Server/internal/dao"
	"Ba-Server/internal/model/entity"
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"
)

type (
	sLogsOption struct{}
)

func init() {
	service.RegisterLogsOption(New())
}

func New() service.ILogsOption {
	return &sLogsOption{}
}

func (s sLogsOption) GetMonLogsOptionPage(ctx context.Context, req *v1.GetMonOptionLogsPageReq) (total int, records []vo.MonLogsOptionVo, err error) {
	// 创建指针切片
	var result []vo.MonLogsOptionVo

	monLogsOptionModel := dao.MonLogsOperation.Ctx(ctx)

	pageQuery := monLogsOptionModel.OmitEmpty().Page(req.Current, req.Size).Where("user_id", req.UserId)
	if req.StartTime != "" || req.EndTime != "" {
		pageQuery.WhereBetween("create_time", req.StartTime, req.EndTime)
	}
	err = pageQuery.ScanAndCount(&result, &total, true)

	if err != nil {
		return 0, nil, err
	}

	// 如果列表为空
	if result == nil {
		records = []vo.MonLogsOptionVo{}
		return
	}

	//添加用户名称
	sysUserModel := dao.SysUser.Ctx(ctx)
	for i, monLogsOption := range result {
		var sysUser entity.SysUser
		err = sysUserModel.Unscoped().Where("id", monLogsOption.UserId).Fields("nick_name").Scan(&sysUser)
		if err != nil {
			result[i].UserName = "未知用户"
			continue
		}
		result[i].UserName = sysUser.NickName
	}

	records = result

	return
}
