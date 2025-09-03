package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func ListResourceWithLimitOffset(ctx context.Context, limit int, offset int) ([]entity.Resource, error) {
	var resources []entity.Resource
	err := mysql.DB.Model(&entity.Resource{}).Order("id DESC").Limit(limit).Offset(offset).Find(&resources).Error
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func ListResourceWithCategoryLimitOffset(ctx context.Context, limit int, offset int, category int) ([]entity.Resource, error) {
	var resources []entity.Resource
	err := mysql.DB.Model(&entity.Resource{}).Order("id DESC").Where("category_id = ?", category).Limit(limit).Offset(offset).Find(&resources).Error
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func ListResourceWithUnitLimitOffset(ctx context.Context, limit int, offset int, unit int) ([]entity.Resource, error) {
	var resources []entity.Resource
	err := mysql.DB.Model(&entity.Resource{}).Order("id DESC").Where("unit_id = ?", unit).Limit(limit).Offset(offset).Find(&resources).Error
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func CreateResource(ctx context.Context, name string, categoryId int, description string, filepath string, coverPath string, unitId uint) error {
	resource := entity.Resource{
		Name:        name,
		CategoryID:  categoryId,
		Description: description,
		FilePath:    filepath,
		CoverPath:   coverPath,
		UnitId:      unitId,
	}
	err := mysql.DB.Create(&resource).Error
	if err != nil {
		return err
	}
	return nil
}
