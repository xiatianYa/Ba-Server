package file

import (
	"Ba-Server/internal/model/vo"
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/file/v1"
)

func (c *ControllerV1) FileUpload(ctx context.Context, req *v1.FileUploadReq) (res *v1.FileUploadRes, err error) {
	if req.File == nil {
		return nil, gerror.New("上传文件不存在!")
	}
	fileId, fileUrl, err := service.File().UploadFile(ctx, req)
	if err != nil {
		return nil, err
	}
	fileVo := vo.MonLogsFileVo{
		Id:       fileId,
		FileName: req.File.Filename,
		FileUrl:  fileUrl,
	}
	return (*v1.FileUploadRes)(&fileVo), nil
}
