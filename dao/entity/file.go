package entity

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Name         string `json:"name" gorm:"not null"`
	CategoryID   uint   `json:"category_id" gorm:"not null"`
	Description  string `json:"description,omitempty"`
	ResourcePath string `json:"resource_path"`
	CoverPath    string `json:"cover_path"`
}
