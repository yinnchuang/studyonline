package entity

import (
	"time"

	"gorm.io/gorm"
)

type Homework struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	FilePath    string    `json:"file_path"`
	ExpireTime  time.Time `json:"expire_time"`
}
