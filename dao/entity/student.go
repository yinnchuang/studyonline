package entity

import (
	"time"
)

type Student struct {
	ID              uint `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            string `json:"name" gorm:"not null"`
	Username        string `json:"username" gorm:"not null;unique"`
	Password        string `json:"password" gorm:"not null"`
	Department      string `json:"department" gorm:"not null"`
	Email           string `json:"email"`
	CommentCount    int    `json:"comment_count" gorm:"default:0"`    // 评论数
	LikeCount       int    `json:"like_count" gorm:"default:0"`       // 点赞数
	BeCommentedCount int   `json:"be_commented_count" gorm:"default:0"` // 被评论数
}
