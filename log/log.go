package log

import (
	"encoding/csv"
	"log"
	"os"
)

var (
	DownloadLogger *csv.Writer
)

func Init() {
	// 打开文件，使用追加模式，如果文件不存在则创建
	logFile, err := os.OpenFile("download.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println("无法打开日志文件: %v", err)
	}

	// 创建CSV写入器
	DownloadLogger = csv.NewWriter(logFile)

	// 检查文件是否为空，如果为空则写入表头
	fileInfo, err := logFile.Stat()
	if err != nil {
		log.Println("获取文件信息失败: %v", err)
	}

	// 如果文件为空，写入CSV表头
	if fileInfo.Size() == 0 {
		headers := []string{"时间", "下载人", "部门", "下载数据集"}
		if err := DownloadLogger.Write(headers); err != nil {
			log.Println("写入CSV表头失败: %v", err)
		}
		DownloadLogger.Flush() // 确保表头被写入
	}
}
