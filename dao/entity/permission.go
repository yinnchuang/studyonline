package entity

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	UserID    uint `json:"user_id" gorm:"not null"`
	Identity  int  `json:"identity" gorm:"not null"`
	DatasetId uint `json:"dataset_id" gorm:"not null"`
	TeacherId uint `json:"teacher_id" gorm:"not null"`
}

type Request struct {
	gorm.Model
	UserID    uint   `json:"user_id" gorm:"not null"`
	Identity  int    `json:"identity" gorm:"not null"`
	DatasetId uint   `json:"dataset_id" gorm:"not null"`
	Reason    string `json:"reason" gorm:"not null"`
	TeacherId uint   `json:"teacher_id" gorm:"not null"`
	Status    int    `json:"status" gorm:"not null"` // 0代表待审核，1代表同意，-1代表拒绝
}
