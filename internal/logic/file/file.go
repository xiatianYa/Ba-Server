package file

import (
	v1 "Ba-Server/api/file/v1"
	"Ba-Server/internal/service"
	"Ba-Server/utility/file"
	"context"
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

func (s sFile) UploadFile(ctx context.Context, req *v1.FileUploadReq) (res string, err error) {
	sysFile := req.File
	//校验文件名称长度 和 文件格式是否符合规定
	fileName := sysFile.Filename
	err = file.ValidateFileNameLength(fileName)
	if err != nil {
		return "", err
	}
	err = file.ValidateFileFormat(fileName)
	if err != nil {
		return "", err
	}
	//校验文件大小是否符合规定
	fileSize := sysFile.Size
	err = file.ValidateFileSize(fileSize)
	if err != nil {
		return "", err
	}
	//上传文件目录
	filePath, err := file.UploadFile(sysFile)
	if err != nil {
		return "", err
	}
	return filePath, nil
}
