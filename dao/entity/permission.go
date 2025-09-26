package entity

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	UserID    uint `json:"user_id" gorm:"not null"`
	Identity  int  `json:"identity" gorm:"not null"`
	DatasetId uint `json:"dataset_id" gorm:"not null"`
	TeacherId uint `json:"teacher_id" gorm:"not null"`
}
