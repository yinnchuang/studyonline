package entity

import "gorm.io/gorm"

type DownloadLog struct {
	gorm.Model
	Downloader           string `gorm:"type:varchar(255);not null"`
	DownloaderDepartment string `gorm:"type:varchar(255);not null"`
	DatasetName          string `gorm:"type:varchar(255);not null"`
}
