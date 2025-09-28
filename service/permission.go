package service

import (
	"context"
	"fmt"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
	"studyonline/dao/redis"

	"gorm.io/gorm"
)

func ListNeedAuthPermission(ctx context.Context, datasetId uint) ([]string, error) {
	key := fmt.Sprintf("permission:%d", datasetId)
	set, err := redis.RDB.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return set, nil
}

func SetNeedAuthPermission(ctx context.Context, datasetId uint, userId uint, identity int) error {
	key := fmt.Sprintf("permission:%d", datasetId)
	value := fmt.Sprintf("%d_%d", userId, identity)
	_, err := redis.RDB.SAdd(ctx, key, value).Result()
	if err != nil {
		return err
	}
	return nil
}

func IfUserHasDatasetPermission(userId uint, identity int, datasetId uint) bool {
	var count int64 = 0
	mysql.DB.Model(&entity.Permission{}).Where("user_id = ?", userId).Where("dataset_id = ?", datasetId).Where("identity = ?", identity).Count(&count)
	return count > 0
}

func CreatePermission(ctx context.Context, userId uint, identity int, datasetId uint, teacherId uint) error {
	permission := entity.Permission{
		UserID:    userId,
		DatasetId: datasetId,
		Identity:  identity,
		TeacherId: teacherId,
	}
	err := mysql.DB.Create(&permission).Error
	if err != nil {
		return err
	}
	key := fmt.Sprintf("permission:%d", datasetId)
	value := fmt.Sprintf("%d_%d", userId, identity)
	_, err = redis.RDB.SRem(ctx, key, value).Result()
	if err != nil {
		return err
	}
	return nil
}

func CreateRequest(ctx context.Context, datasetId uint, userId uint, identity int, reason string, teacherId uint) error {
	request := entity.Request{
		UserID:    userId,
		Identity:  identity,
		DatasetId: datasetId,
		Reason:    reason,
		TeacherId: teacherId,
		Status:    0,
	}
	err := mysql.DB.Model(&entity.Request{}).Create(&request).Error
	if err != nil {
		return err
	}
	return nil
}

func ListRequestByUserId(ctx context.Context, userId uint, identity int) ([]entity.Request, error) {
	var requests []entity.Request
	err := mysql.DB.Model(&entity.Request{}).Where("user_id = ? AND identity = ?", userId, identity).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func ListRequestByTeacherId(ctx context.Context, teacherId uint) ([]entity.Request, error) {
	var requests []entity.Request
	err := mysql.DB.Model(&entity.Request{}).Where("teacher_id = ?", teacherId).Find(&requests).Error
	if err != nil {
		return nil, err
	}
	return requests, nil
}

func AgreeRequest(ctx context.Context, requestId uint, userId uint, identity int, datasetId uint, teacherId uint) error {
	err := mysql.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&entity.Request{}).Where("id = ?", requestId).Update("status", 1).Error
		if err != nil {
			return err
		}
		permission := entity.Permission{
			UserID:    userId,
			Identity:  identity,
			DatasetId: datasetId,
			TeacherId: teacherId,
		}
		err = tx.Model(&entity.Permission{}).FirstOrCreate(&permission).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func DisagreeRequest(ctx context.Context, requestId uint, userId uint, identity int, datasetId uint, teacherId uint) error {
	err := mysql.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&entity.Request{}).Where("id = ?", requestId).Update("status", -1).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func GetRequestById(ctx context.Context, requestId uint) (*entity.Request, error) {
	var request entity.Request
	err := mysql.DB.Model(&entity.Request{}).Where("id = ?", requestId).First(&request).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}
