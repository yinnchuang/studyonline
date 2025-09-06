package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func GetAllDiscusses(c context.Context) ([]entity.Discuss, error) {
	var discusses []entity.Discuss
	err := mysql.DB.Model(&entity.Discuss{}).Find(&discusses).Error
	if err != nil {
		return nil, err
	}
	return discusses, nil
}

func CreateDiscuss(c context.Context, discuss entity.Discuss) error {
	err := mysql.DB.Model(&entity.Discuss{}).Create(&discuss).Error
	return err
}

func RemoveDiscuss(c context.Context, id uint) error {
	err := mysql.DB.Model(&entity.Discuss{}).Delete(&entity.Discuss{}, id).Error
	return err
}
