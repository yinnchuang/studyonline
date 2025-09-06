package service

import (
	"studyonline/constant"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

type GetUserInfoVO struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

func GetUserInfo(userId uint, identity int) (*GetUserInfoVO, error) {
	if identity == constant.StudentIdentity {
		return GetStudentInfo(userId)
	} else if identity == constant.TeacherIdentity {
		return GetTeacherInfo(userId)
	}
	return nil, nil
}

func GetStudentInfo(id uint) (*GetUserInfoVO, error) {
	var result GetUserInfoVO
	err := mysql.DB.Model(&entity.Student{}).Where("id = ?").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetTeacherInfo(id uint) (*GetUserInfoVO, error) {
	var result GetUserInfoVO
	err := mysql.DB.Model(&entity.Teacher{}).Where("id = ?").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
