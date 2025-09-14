package entity

import "gorm.io/gorm"

type Score struct {
	gorm.Model
	StudentId  uint `json:"student_id" gorm:"not null"`
	UsualScore int  `json:"usual_score" gorm:"not null"`
	ExamScore  int  `json:"exam_score" gorm:"not null"`
	FinalScore int  `json:"final_score" gorm:"not null"`
}
