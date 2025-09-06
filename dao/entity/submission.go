package entity

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	StudentId   uint   `json:"student_id" gorm:"not null"`
	HomeworkId  uint   `json:"homework_id" gorm:"not null"`
	FilePath    string `json:"file_path"`
	Description string `json:"description"`
}
