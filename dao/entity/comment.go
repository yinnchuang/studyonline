package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserId    uint   `json:"user_id"`
	Identity  int    `json:"identity"`
	DiscussId uint   `json:"discuss_id"`
	Comment   string `json:"comment"`
}
