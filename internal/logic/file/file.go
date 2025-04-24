package file

import (
	v1 "Ba-Server/api/file/v1"
	"Ba-Server/internal/consts"
	"Ba-Server/internal/dao"
	"Ba-Server/internal/model/entity"
	"Ba-Server/internal/service"
	"Ba-Server/utility/file"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"strconv"
)

type (
	sFile struct{}
)

func init() {
	service.RegisterFile(New())
}

func New() service.IFile {
	return &sFile{}
}

func (s sFile) UploadFile(ctx context.Context, req *v1.FileUploadReq) (fileId *int64, fileUrl *string, err error) {
	//上传用户ID
	userId := ctx.Value(consts.UserId)
	if userId == nil {
		return nil, nil, gerror.New("用户信息获取失败,请重新登陆!")
	}
	var parseUserId int64
	err = gconv.Scan(userId, &parseUserId)
	MonLogsFileModel := dao.MonLogsFile.Ctx(ctx)
	sysFile := req.File
	//校验文件名称长度 和 文件格式是否符合规定
	fileName := sysFile.Filename
	err = file.ValidateFileNameLength(fileName)
	err = file.ValidateFileFormat(fileName)
	//校验文件大小是否符合规定
	fileSize := sysFile.Size
	err = file.ValidateFileSize(fileSize)
	//获取文件类型
	fileType := file.GetFileType(fileName)
	//上传文件目录
	fileUrl, filePath, err := file.UploadFile(sysFile)
	//如果文件上传保存 则记录日志 成功也纪录
	MonLogFile := entity.MonLogsFile{
		UserId:   parseUserId,
		FilePath: *filePath,
		FileUrl:  *fileUrl,
		FileType: fileType,
		FileSize: strconv.FormatInt(fileSize, 10),
	}
	if err != nil {
		MonLogFile.Status = consts.ZERO
		MonLogFile.ErrorMsg = err.Error()
		return nil, nil, err
	} else {
		MonLogFile.Status = consts.ONE
	}
	id, err := MonLogsFileModel.InsertAndGetId(MonLogFile)
	if err != nil {
		return nil, nil, gerror.New("文件上传日志保存失败" + err.Error())
	}
	return &id, fileUrl, nil
}
