package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func AddLog(ctx context.Context, downloader string, downloaderDepartment string, datasetName string) {
	mysql.DB.Model(&entity.DownloadLog{}).Create(&entity.DownloadLog{
		Downloader:           downloader,
		DownloaderDepartment: downloaderDepartment,
		DatasetName:          datasetName,
	})
}
