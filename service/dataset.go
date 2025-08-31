package service

import (
	"context"
	"path/filepath"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func ListDatasetWithLimit(ctx context.Context, limit int, offset int) ([]entity.Dataset, error) {
	datasets := []entity.Dataset{}
	err := mysql.DB.Model(&entity.Dataset{}).Order("id DESC").Limit(limit).Offset(offset).Find(&datasets).Error
	if err != nil {
		return nil, err
	}
	return datasets, nil
}

func ListDatasetWithCategory(ctx context.Context, limit int, offset int, category int) ([]entity.Dataset, error) {
	datasets := []entity.Dataset{}
	err := mysql.DB.Model(&entity.Dataset{}).Order("id DESC").Where("category_id = ?", category).Limit(limit).Offset(offset).Find(&datasets).Error
	if err != nil {
		return nil, err
	}
	return datasets, nil
}

func CreateDataset(ctx context.Context, name string, categoryId int, description string, filePath string, coverPath string) error {
	resourceAbsPath, _ := filepath.Abs(filePath)
	coverAbsPath, _ := filepath.Abs(coverPath)
	dataset := entity.Dataset{
		Name:        name,
		CategoryID:  categoryId,
		Description: description,
		FilePath:    resourceAbsPath,
		CoverPath:   coverAbsPath,
	}
	err := mysql.DB.Create(&dataset).Error
	if err != nil {
		return err
	}
	return nil
}
