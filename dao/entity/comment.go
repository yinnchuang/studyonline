package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserId    uint   `json:"user_id"`
	DiscussId uint   `json:"discuss_id"`
	Comment   string `json:"comment"`
	OwnerName string `json:"owner_name"`
}
