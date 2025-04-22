package v1

import (
	"Ba-Server/internal/model/vo"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/upload" method:"post" mime:"multipart/form-data" tags:"上传" summary:"上传文件"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"上传的文件"`
}

type FileUploadRes vo.MonLogsFileVo
