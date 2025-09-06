package entity

import "gorm.io/gorm"

type Discuss struct {
	gorm.Model
	UserId      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	OwnerName   string `json:"owner_name"`
}
