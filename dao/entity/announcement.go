package entity

import "gorm.io/gorm"

type Announcement struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}
