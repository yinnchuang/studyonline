package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"

	"gorm.io/gorm"
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

func UpdateResource(ctx context.Context, resourceId uint, name string, categoryId int, description string, filepath string, coverPath string, unitIds []uint) error {
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
	err := mysql.DB.Model(&entity.Resource{}).Where("id = ?", resourceId).Updates(resource).Error
	if err != nil {
		return err
	}
	return nil
}

func GetResourceByID(ctx context.Context, id uint) (*entity.Resource, error) {
	var resource entity.Resource
	err := mysql.DB.First(&resource, id).Error
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

func PlusResourceDownloadTime(ctx context.Context, resourceId uint) error {
	return mysql.DB.Model(&entity.Resource{}).Where("id = ?", resourceId).
		Update("download_time", gorm.Expr("download_time + ?", 1)).
		Error
}

func DeleteResource(ctx context.Context, id uint) error {
	return mysql.DB.Delete(&entity.Resource{}, id).Error
}

func SearchResourceByKeyword(ctx context.Context, limit int, offset int, keyword string) ([]entity.Resource, error) {
	var resources []entity.Resource
	// 构建查询条件：名称或描述包含关键词（模糊匹配）
	query := mysql.DB.Preload("Units").
		Where("name LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%").
		Order("id DESC") // 保持与其他列表接口一致的排序（按ID倒序）

	// 处理分页逻辑（复用原有 limit=-1 表示查询全部的逻辑）
	if limit == -1 {
		err := query.Find(&resources).Error
		if err != nil {
			return nil, err
		}
		return resources, nil
	}

	// 带分页的查询
	err := query.Limit(limit).Offset(offset).Find(&resources).Error
	if err != nil {
		return nil, err
	}
	return resources, nil
}

// CountResourceByKeyword 统计符合关键词搜索条件的资源总数
func CountResourceByKeyword(ctx context.Context, keyword string) (int64, error) {
	var count int64
	// 与搜索逻辑保持一致的条件（名称或描述包含关键词）
	err := mysql.DB.Model(&entity.Resource{}).
		Where("name LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%").
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
