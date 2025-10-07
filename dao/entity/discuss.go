package entity

import "gorm.io/gorm"

type Discuss struct {
	gorm.Model
	UserId      uint   `json:"user_id"`
	Identity    int    `json:"identity"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
