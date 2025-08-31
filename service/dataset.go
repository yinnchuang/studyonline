package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func ListDatasetWithLimitOffset(ctx context.Context, limit int, offset int) ([]entity.Dataset, error) {
	var datasets []entity.Dataset
	err := mysql.DB.Model(&entity.Dataset{}).Order("id DESC").Limit(limit).Offset(offset).Find(&datasets).Error
	if err != nil {
		return nil, err
	}
	return datasets, nil
}

func ListDatasetWithCategoryLimitOffset(ctx context.Context, limit int, offset int, category int) ([]entity.Dataset, error) {
	var datasets []entity.Dataset
	err := mysql.DB.Model(&entity.Dataset{}).Order("id DESC").Where("category_id = ?", category).Limit(limit).Offset(offset).Find(&datasets).Error
	if err != nil {
		return nil, err
	}
	return datasets, nil
}

func ListDatasetWithUnit(ctx context.Context, unit int) ([]entity.Dataset, error) {
	var datasets []entity.Dataset
	err := mysql.DB.Model(&entity.Dataset{}).Order("id DESC").Where("unit_id = ?", unit).Find(&datasets).Error
	if err != nil {
		return nil, err
	}
	return datasets, nil
}

func CreateDataset(ctx context.Context, name string, categoryId int, description string, filePath string, coverPath string, unitId uint) error {
	dataset := entity.Dataset{
		Name:        name,
		CategoryID:  categoryId,
		Description: description,
		FilePath:    filePath,
		CoverPath:   coverPath,
		UnitId:      unitId,
	}
	err := mysql.DB.Create(&dataset).Error
	if err != nil {
		return err
	}
	return nil
}
