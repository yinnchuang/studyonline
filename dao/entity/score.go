package entity

import "gorm.io/gorm"

type Score struct {
	gorm.Model
	StudentId     uint `json:"student_id" gorm:"not null"`
	UsualScore    int  `json:"usual_score" gorm:"not null"`
	ExamScore     int  `json:"exam_score" gorm:"not null"`
	FinalScore    int  `json:"final_score"` // 保留字段兼容历史数据，但不再由前端传入
	ActivityScore int  `json:"activity_score" gorm:"default:0"` // 活跃度分数
}
