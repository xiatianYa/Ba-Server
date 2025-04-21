package vo

// MonLogsFileVo is the golang structure for table mon_logs_file.
type MonLogsFileVo struct {
	Id       interface{} `json:"id"`   // 主键ID
	FileName string      `json:"name"` //文件名称
	FileUrl  interface{} `json:"url"`  // 文件网络路径
}
