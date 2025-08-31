package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func GetAllHomework(ctx context.Context) ([]entity.Homework, error) {
	var homeworks []entity.Homework
	err := mysql.DB.Model(&entity.Homework{}).Find(&homeworks).Error
	if err != nil {
		return nil, err
	}
	return homeworks, nil
}

func CreateHomework(ctx context.Context, homework entity.Homework) error {
	err := mysql.DB.Model(&entity.Homework{}).Create(&homework).Error
	if err != nil {
		return err
	}
	return nil
}

func RemoveHomework(ctx context.Context, id uint) error {
	err := mysql.DB.Where("id = ?", id).Delete(&entity.Homework{}).Error
	if err != nil {
		return err
	}
	return nil
}
