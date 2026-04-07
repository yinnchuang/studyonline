package entity

import "gorm.io/gorm"

type LessonPlanDiscussLike struct {
	gorm.Model
	DiscussID uint `json:"discuss_id"` // 评论id
	UserId    uint `json:"user_id"`    // 用户id
	Identity  int  `json:"identity"`   // 用户身份
}
