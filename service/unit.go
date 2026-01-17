package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func GetAllUnit(c context.Context) ([]entity.Unit, error) {
	var units []entity.Unit
	err := mysql.DB.Model(&entity.Unit{}).Find(&units).Error
	if err != nil {
		return nil, err
	}
	return units, nil
}

func RemoveUnit(c context.Context, unitId uint) error {
	err := mysql.DB.Model(&entity.Unit{}).Where("id = ?", unitId).Delete(&entity.Unit{}).Error
	return err
}

func GetUnitById(c context.Context, unitId uint) (*entity.Unit, error) {
	var unit entity.Unit
	err := mysql.DB.Model(&entity.Unit{}).Where("id = ?", unitId).First(&unit).Error
	if err != nil {
		return nil, err
	}
	return &unit, nil
}

func GetSonUnit(c context.Context, fatherUnitId uint) ([]entity.Unit, error) {
	var units []entity.Unit
	err := mysql.DB.Model(&entity.Unit{}).Where("father_unit_id = ?", fatherUnitId).Find(&units).Error
	if err != nil {
		return nil, err
	}
	return units, nil
}

func CreateUnit(c context.Context, unit entity.Unit) error {
	err := mysql.DB.Create(&unit).Error
	return err
}
