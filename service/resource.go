package service

import (
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func RandomListResource() ([]entity.Resource, error) {
	resources := []entity.Resource{}
	err := mysql.DB.Model(&entity.Resource{}).Order("RAND()").Limit(10).Find(&resources).Error
	if err != nil {
		return nil, err
	}
	return resources, nil
}
