package entity

import (
	"time"
)

type Teacher struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string `json:"name" gorm:"not null"`
	Username   string `json:"username" gorm:"not null;unique"`
	Password   string `json:"password" gorm:"not null"`
	Department string `json:"department" gorm:"not null"`
	Email      string `json:"email" gorm:"not null;unique"`
}
