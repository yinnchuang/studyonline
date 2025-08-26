package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func ListDatasetWithLimit(ctx context.Context, limit int) ([]entity.Dataset, error) {
	datasets := []entity.Dataset{}
	err := mysql.DB.Model(&entity.Dataset{}).Order("id DESC").Limit(limit).Find(&datasets).Error
	if err != nil {
		return nil, err
	}
	return datasets, nil
}
