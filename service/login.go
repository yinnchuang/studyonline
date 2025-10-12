package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"studyonline/constant"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
	"studyonline/dao/redis"
	"studyonline/util"
)

func Login(ctx context.Context, username string, password string, identity int) (bool, string, error) {
	var idFromDB uint
	var passwordFromDB string
	if identity == constant.StudentIdentity { // 学生登录
		stu := entity.Student{}
		err := mysql.DB.Model(&entity.Student{}).Where("username = ?", username).Find(&stu).Error
		if err != nil {
			return false, "", err
		}
		idFromDB = stu.ID
		passwordFromDB = stu.Password
	} else if identity == constant.TeacherIdentity { // 教师登录
		tea := entity.Teacher{}
		err := mysql.DB.Model(&entity.Teacher{}).Where("username = ?", username).Find(&tea).Error
		if err != nil {
			return false, "", err
		}
		idFromDB = tea.ID
		passwordFromDB = tea.Password
	} else if identity == constant.AdminIdentity { // 管理员登录
		adm := entity.Admin{}
		err := mysql.DB.Model(&entity.Admin{}).Where("username = ?", username).Find(&adm).Error
		if err != nil {
			return false, "", err
		}
		idFromDB = adm.ID
		passwordFromDB = adm.Password
	} else {
		return false, "", errors.New("登录失败")
	}
	login := util.ComparePwd(passwordFromDB, password)
	//if identity == constant.AdminIdentity {
	//	login = password == passwordFromDB
	//}
	if login {
		cacheKey := util.GenerateToken()
		cacheValue := fmt.Sprintf("%v_%v", idFromDB, identity)
		redis.RDB.Set(ctx, cacheKey, cacheValue, time.Hour*6)
		return true, cacheKey, nil
	} else {
		return false, "", nil
	}
}
