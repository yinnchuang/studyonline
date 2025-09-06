package entity

import "gorm.io/gorm"

type Score struct {
	gorm.Model
	StudentId  uint `json:"student_id" gorm:"not null"`
	TeacherId  uint `json:"teacher_id" gorm:"not null"`
	HomeworkId uint `json:"homework_id" gorm:"not null"`
	Score      int  `json:"score" gorm:"not null"`
}
