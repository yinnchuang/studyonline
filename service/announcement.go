package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func GetAllAnnouncement(c context.Context) ([]entity.Announcement, error) {
	var announcements []entity.Announcement
	err := mysql.DB.Model(&entity.Announcement{}).Find(&announcements).Error
	if err != nil {
		return nil, err
	}
	return announcements, nil
}

func CreateAnnouncement(c context.Context, announcement entity.Announcement) error {
	err := mysql.DB.Model(&entity.Announcement{}).Create(&announcement).Error
	return err
}
