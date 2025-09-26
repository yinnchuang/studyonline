package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"

	"github.com/gin-gonic/gin"
)

func ListDatasetByTeacherId(c *gin.Context, teacherId uint) ([]entity.Dataset, error) {
	var datasets []entity.Dataset
	err := mysql.DB.Model(&entity.Dataset{}).Where("teacher_id = ?", teacherId).Find(&datasets).Error
	if err != nil {
		return nil, err
	}
	return datasets, nil
}

func ListDatasetWithLimitOffset(ctx context.Context, limit int, offset int) ([]entity.Dataset, error) {
	var datasets []entity.Dataset
	err := mysql.DB.Model(&entity.Dataset{}).Order("id DESC").Limit(limit).Offset(offset).Find(&datasets).Error
	if err != nil {
		return nil, err
	}
	return datasets, nil
}

func CountDataset(ctx context.Context) (int64, error) {
	var count int64
	err := mysql.DB.Model(&entity.Dataset{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func ListDatasetWithCategoryLimitOffset(ctx context.Context, limit int, offset int, category int) ([]entity.Dataset, error) {
	var datasets []entity.Dataset
	err := mysql.DB.Model(&entity.Dataset{}).Order("id DESC").Where("category_id = ?", category).Limit(limit).Offset(offset).Find(&datasets).Error
	if err != nil {
		return nil, err
	}
	return datasets, nil
}

func CountDatasetWithCategory(ctx context.Context, category int) (int64, error) {
	var count int64
	err := mysql.DB.Model(&entity.Dataset{}).Where("category_id = ?", category).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CreateDataset(ctx context.Context, name string, categoryId int, description string, filePath string, coverPath string, scale string, teacherId uint, private bool) (*entity.Dataset, error) {
	dataset := entity.Dataset{
		Name:        name,
		CategoryID:  categoryId,
		Description: description,
		Scale:       scale,
		FilePath:    filePath,
		CoverPath:   coverPath,
		TeacherId:   teacherId,
		Private:     private,
	}
	err := mysql.DB.Create(&dataset).Error
	if err != nil {
		return nil, err
	}
	return &dataset, nil
}

func GetDatasetByID(ctx context.Context, id uint) (*entity.Dataset, error) {
	var dataset entity.Dataset
	err := mysql.DB.First(&dataset, id).Error
	if err != nil {
		return nil, err
	}
	return &dataset, nil
}
