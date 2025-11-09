package service

import (
	"context"
	"studyonline/constant"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
	"studyonline/dao/redis"
	"studyonline/util"
	"time"
)

type GetUserInfoVO struct {
	Name       string `json:"name"`
	Username   string `json:"username"`
	Department string `json:"department"`
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
	res := mysql.DB.Model(&entity.Student{}).Where("id = ?", id).Find(&result)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return &result, nil
}

func GetStudentInfoByUsername(username string) (*entity.Student, error) {
	var result entity.Student
	res := mysql.DB.Model(&entity.Student{}).Where("username = ?", username).Find(&result)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return &result, nil
}

func GetTeacherInfo(id uint) (*GetUserInfoVO, error) {
	var result GetUserInfoVO
	res := mysql.DB.Model(&entity.Teacher{}).Where("id = ?", id).Find(&result)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return &result, nil
}

func GetTeacherInfoByUsername(username string) (*entity.Teacher, error) {
	var result entity.Teacher
	res := mysql.DB.Model(&entity.Teacher{}).Where("username = ?", username).Find(&result)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return &result, nil
}

func ChangeStudentPassword(studentId uint, password string) error {
	return mysql.DB.Model(&entity.Student{}).Where("id = ?", studentId).Update("password", password).Error
}

func ChangeTeacherPassword(teacherId uint, password string) error {
	return mysql.DB.Model(&entity.Teacher{}).Where("id = ?", teacherId).Update("password", password).Error
}

func BindStudentEmail(studentId uint, email string) error {
	return mysql.DB.Model(&entity.Student{}).Where("id = ?", studentId).Update("email", email).Error
}

func BindTeacherEmail(teacherId uint, email string) error {
	return mysql.DB.Model(&entity.Teacher{}).Where("id = ?", teacherId).Update("email", email).Error
}

func SendCode2Email(ctx context.Context, email string) error {
	code := util.GenerateCode()
	cacheKey := email
	cacheValue := code
	if err := redis.RDB.Set(ctx, cacheKey, cacheValue, time.Minute*5).Err(); err != nil {
		return err
	}
	err := util.SendCodeToEmail(email, code)
	if err != nil {
		return err
	}
	return nil
}
