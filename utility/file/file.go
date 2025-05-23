package file

import (
	"Ba-Server/internal/consts"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FilePath 基础上传路径
var FilePath = "C:/jinLink/uploadPath"

// BasePath 文件保存路径
var BasePath = "http://127.0.0.1:8080/statics/"

// 允许的文件格式
var allowedExtensions = []string{
	// 图片
	"bmp", "gif", "jpg", "jpeg", "png", "svg",
	// word excel powerpoint
	"doc", "docx", "xls", "xlsx", "ppt", "pptx", "html", "txt",
	// 压缩文件
	"rar", "zip", "gz",
	// 视频格式
	"mp4",
	// 音乐
	"mp3",
	// pdf
	"pdf",
}

// 文件最大长度
var maxLength = 20

// 文件最大200MB
var maxSize int64 = 209715200

// ValidateFileNameLength 校验文件名称长度
func ValidateFileNameLength(fileName string) error {
	if len(fileName) > maxLength {
		return gerror.New("文件名称长度超过限制")
	}
	return nil
}

// ValidateFileFormat 校验文件格式
func ValidateFileFormat(fileName string) error {
	ext := strings.TrimPrefix(filepath.Ext(fileName), ".")
	ext = strings.ToLower(ext)
	isValidFormat := false
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			isValidFormat = true
			break
		}
	}
	if !isValidFormat {
		return gerror.New("不支持的文件格式")
	}
	return nil
}

// ValidateFileSize 校验文件大小
func ValidateFileSize(fileSize int64) error {
	if fileSize > maxSize {
		return gerror.New("文件大小超过限制")
	}
	return nil
}

// GetFileType 获取文件类型，返回对应的数字标识
func GetFileType(filename string) int {
	// 先尝试通过扩展名判断
	ext := filepath.Ext(filename)
	if ext != "" {
		ext = strings.ToLower(strings.TrimPrefix(ext, "."))
		switch ext {
		case "jpg", "jpeg", "png", "gif", "bmp":
			return consts.FILE_IMG
		case "mp4", "avi", "mov", "mkv":
			return consts.FILE_VIDEO
		case "zip", "rar", "7z":
			return consts.FILE_COMPRESS
		case "mp3":
			return consts.FILE_AUDIO
		}
	}
	return consts.FILE_FILE
}

// UploadFile 上传文件
func UploadFile(file *ghttp.UploadFile) (*string, *string, error) {
	// 获取当前日期
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()

	// 构建目标目录
	targetDir := fmt.Sprintf("%d/%02d/%02d/", year, month, day)

	// 创建目标目录
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return nil, nil, fmt.Errorf("failed to create target directory: %w", err)
	}

	// 构建完整的文件路径
	targetPath := filepath.Join(FilePath, targetDir)

	// 保存文件
	filename, err := file.Save(targetPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to save file: %w", err)
	}

	targetUrl := BasePath + targetDir + filename
	targetPath = filepath.Join(targetPath, filename)

	return &targetUrl, &targetPath, nil
}

// RemoveFile 删除文件
func RemoveFile(filePath string) bool {
	err := os.Remove(filePath)
	if err != nil {
		return false
	}
	return true
}

// CreateDirIfNotExists 创建目录
func CreateDirIfNotExists(dir string) error {
	// 获取目录的绝对路径
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	// 检查目录是否存在
	fileInfo, err := os.Stat(absDir)
	if os.IsNotExist(err) {
		// 若不存在则创建目录，权限设置为 0755
		err = os.MkdirAll(absDir, 0755)
		if err != nil {
			return fmt.Errorf("创建目录 %s 失败: %w", absDir, err)
		}
		fmt.Printf("目录 %s 创建成功\n", absDir)
	} else if err != nil {
		return fmt.Errorf("检查目录 %s 时出错: %w", absDir, err)
	} else if !fileInfo.IsDir() {
		return fmt.Errorf("%s 不是一个目录", absDir)
	} else {
		fmt.Printf("目录 %s 已存在\n", absDir)
	}
	return nil
}
