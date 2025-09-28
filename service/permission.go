package service

import (
	"context"
	"fmt"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
	"studyonline/dao/redis"
)

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

func CreateRequest(ctx context.Context, datasetId uint, userId uint, identity int) error {
	
}
