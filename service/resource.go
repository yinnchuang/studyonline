package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func ListResourceWithLimitOffset(ctx context.Context, limit int, offset int) ([]entity.Resource, error) {
	var resources []entity.Resource
	if limit == -1 {
		err := mysql.DB.Preload("Units").Model(&entity.Resource{}).Find(&resources).Error
		if err != nil {
			return nil, err
		}
		return resources, nil
	}

	err := mysql.DB.Preload("Units").Model(&entity.Resource{}).Order("id DESC").Limit(limit).Offset(offset).Find(&resources).Error
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func CountResource(ctx context.Context) (int64, error) {
	var count int64
	err := mysql.DB.Model(&entity.Resource{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func ListResourceWithCategoryLimitOffset(ctx context.Context, limit int, offset int, category int) ([]entity.Resource, error) {
	var resources []entity.Resource
	if limit == -1 {
		err := mysql.DB.Preload("Units").Model(&entity.Resource{}).Order("id DESC").Where("category_id = ?", category).Find(&resources).Error
		if err != nil {
			return nil, err
		}
		return resources, nil
	}

	err := mysql.DB.Preload("Units").Model(&entity.Resource{}).Order("id DESC").Where("category_id = ?", category).Limit(limit).Offset(offset).Find(&resources).Error
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func CountResourceWithCategory(ctx context.Context, category int) (int64, error) {
	var count int64
	err := mysql.DB.Model(&entity.Resource{}).Where("category_id = ?", category).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func ListResourceWithUnitLimitOffset(ctx context.Context, limit int, offset int, unitIds []uint) ([]entity.Resource, error) {
	var resources []entity.Resource
	if limit == -1 {
		err := mysql.DB.Preload("Units").
			Joins("JOIN resource_units ru ON ru.resource_id = resources.id").
			Where("ru.unit_id IN ?", unitIds).
			Distinct().
			Find(&resources).Error
		if err != nil {
			return nil, err
		}
		return resources, nil
	}
	err := mysql.DB.Preload("Units").
		Joins("JOIN resource_units ru ON ru.resource_id = resources.id").
		Where("ru.unit_id IN ?", unitIds).
		Distinct().
		Limit(limit).Offset(offset).Find(&resources).Error
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func CreateResource(ctx context.Context, name string, categoryId int, description string, filepath string, coverPath string, unitIds []uint) error {
	units := make([]entity.Unit, len(unitIds))
	for i, id := range unitIds {
		var unit entity.Unit
		unit.ID = id
		units[i] = unit
	}
	resource := entity.Resource{
		Name:        name,
		CategoryID:  categoryId,
		Description: description,
		FilePath:    filepath,
		CoverPath:   coverPath,
		Units:       units,
	}
	err := mysql.DB.Create(&resource).Error
	if err != nil {
		return err
	}
	return nil
}
