package service

import (
	"context"
	"path/filepath"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func ListResourceWithLimit(ctx context.Context, limit int) ([]entity.Resource, error) {
	resources := []entity.Resource{}
	err := mysql.DB.Model(&entity.Resource{}).Order("id DESC").Limit(limit).Find(&resources).Error
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func CreateResource(ctx context.Context, name string, categoryId int, description string, resourcePath string, coverPath string) error {
	resourceAbsPath, _ := filepath.Abs(resourcePath)
	coverAbsPath, _ := filepath.Abs(coverPath)
	resource := entity.Resource{
		Name:         name,
		CategoryID:   categoryId,
		Description:  description,
		ResourcePath: resourceAbsPath,
		CoverPath:    coverAbsPath,
	}
	err := mysql.DB.Create(&resource).Error
	if err != nil {
		return err
	}
	return nil
}
