package entity

import "gorm.io/gorm"

type Dataset struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	CategoryID  int    `json:"category_id" gorm:"not null"`
	Description string `json:"description,omitempty"`
	Scale       string `json:"scale"`
	FilePath    string `json:"file_path"`
	CoverPath   string `json:"cover_path"`
}
