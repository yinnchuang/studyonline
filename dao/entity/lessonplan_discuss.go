package entity

import "gorm.io/gorm"

type LessonPlanDiscuss struct {
	gorm.Model
	LessonPlanID uint   `json:"lesson_plan_id"` // 绑定某个知识点
	UserId       uint   `json:"user_id"`
	Identity     int    `json:"identity"`
	FatherID     uint   `json:"father_id"` // 父讨论id
	Content      string `json:"content"`
}
