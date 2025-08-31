package entity

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	CategoryID  int    `json:"category_id" gorm:"not null"`
	Description string `json:"description,omitempty"`
	FilePath    string `json:"file_path"`
	CoverPath   string `json:"cover_path"`
	UnitId      uint   `json:"unit_id"`
}
