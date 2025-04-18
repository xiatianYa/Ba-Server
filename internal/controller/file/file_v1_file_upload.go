package file

import (
	"Ba-Server/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"Ba-Server/api/file/v1"
)

func (c *ControllerV1) FileUpload(ctx context.Context, req *v1.FileUploadReq) (res *v1.FileUploadRes, err error) {
	if req.File == nil {
		return nil, gerror.New("请选择需要上传的文件!")
	}
	filePath, err := service.File().UploadFile(ctx, req)
	if err != nil {
		return nil, err
	}
	return (*v1.FileUploadRes)(&filePath), nil
}
