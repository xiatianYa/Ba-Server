package logsFile

import (
	v1 "Ba-Server/api/logsFile/v1"
	"Ba-Server/internal/dao"
	"Ba-Server/internal/model/entity"
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"Ba-Server/utility/file"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
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
	pageQuery := monLogsFileModel.Page(req.Current, req.Size).OmitEmpty().Where("user_id", req.UserId)
	if req.StartTime != "" || req.EndTime != "" {
		pageQuery.WhereBetween("create_time", req.StartTime, req.EndTime)
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

	//添加用户名称
	sysUserModel := dao.SysUser.Ctx(ctx)
	for i, monLogsFile := range result {
		var sysUser entity.SysUser
		err = sysUserModel.Unscoped().Where("id", monLogsFile.UserId).Fields("nick_name").Scan(&sysUser)
		if err != nil {
			result[i].UserName = "未知用户"
			continue
		}
		result[i].UserName = sysUser.NickName
	}

	records = result

	return
}

// RemoveLogsFile 删除单个文件日志
func (s sLogsFile) RemoveLogsFile(ctx context.Context, req *v1.RemoveLogsFileReq) (bool, error) {
	var monLogsFile *entity.MonLogsFile
	monLogsFileModel := dao.MonLogsFile.Ctx(ctx)
	//查询当前文件日志 删除其保留目录
	err := monLogsFileModel.Where("id", req.Id).Scan(&monLogsFile)
	if err != nil {
		return false, err
	}
	if monLogsFile == nil {
		return false, gerror.New("未搜索到文件日志")
	}
	//删除文件
	_, err = monLogsFileModel.Where("id", req.Id).Delete()
	//删除文件所存储位置
	file.RemoveFile(monLogsFile.FilePath)
	if err != nil {
		return false, err
	}
	return true, nil
}
