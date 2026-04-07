package service

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"studyonline/config"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetLessonPlanDiscussByID(c *gin.Context, id uint) (*entity.LessonPlanDiscuss, error) {
	var discuss entity.LessonPlanDiscuss
	err := mysql.DB.Model(&entity.LessonPlanDiscuss{}).Where("id = ?", id).Find(&discuss).Error
	if err != nil {
		return nil, err
	}
	return &discuss, nil
}

func GetLessonPlanDiscussByLessonPlanID(c *gin.Context, id uint) ([]entity.LessonPlanDiscuss, error) {
	var discusses []entity.LessonPlanDiscuss
	err := mysql.DB.Model(&entity.LessonPlanDiscuss{}).Where("lesson_plan_id = ?", id).Find(&discusses).Error
	if err != nil {
		return nil, err
	}
	return discusses, nil
}

func CreateLessonPlanDiscuss(c context.Context, discuss entity.LessonPlanDiscuss) error {
	err := mysql.DB.Model(&entity.LessonPlanDiscuss{}).Create(&discuss).Error
	return err
}

func RemoveLessonPlanDiscussByID(c context.Context, id uint) error {
	err := mysql.DB.Model(&entity.LessonPlanDiscuss{}).Delete(&entity.LessonPlanDiscuss{}, id).Error
	return err
}

func RemoveLessonPlanDiscussByFatherID(c context.Context, id uint) error {
	err := mysql.DB.Where("father_id = ?", id).Delete(&entity.LessonPlanDiscuss{}).Error
	return err
}

func LikeLessonPlanDiscuss(c context.Context, discussID, userID uint, identity int) error {
	// 检查是否已经点赞
	var like entity.LessonPlanDiscussLike
	err := mysql.DB.Where("discuss_id = ? AND user_id = ? AND identity = ?", discussID, userID, identity).First(&like).Error
	if err == nil {
		// 已经点赞，不需要重复操作
		return nil
	}
	
	// 开启事务
	tx := mysql.DB.Begin()
	
	// 创建点赞记录
	newLike := entity.LessonPlanDiscussLike{
		DiscussID: discussID,
		UserId:    userID,
		Identity:  identity,
	}
	if err := tx.Create(&newLike).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	// 更新评论的点赞数
	if err := tx.Model(&entity.LessonPlanDiscuss{}).Where("id = ?", discussID).Update("like_count", gorm.Expr("like_count + ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	return tx.Commit().Error
}

func GetLessonPlanDiscussLikeStatus(c context.Context, discussID, userID uint, identity int) (bool, error) {
	var like entity.LessonPlanDiscussLike
	err := mysql.DB.Where("discuss_id = ? AND user_id = ? AND identity = ?", discussID, userID, identity).First(&like).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func GetLessonPlanDiscussSummary(c context.Context, content string) (string, error) {
	apiKey := config.AppConfig.Deepseek.APIKey
	url := config.AppConfig.Deepseek.URL
	
	reqBody := map[string]interface{}{
		"model": config.AppConfig.Deepseek.Model,
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "你是一个评论整理专家，擅长从大量评论中提取关键信息，总结主要观点和意见。",
			},
			{
				"role":    "user",
				"content": "请帮我整理以下评论，提取主要观点和意见，生成一份简洁明了的总结：\n" + content,
			},
		},
		"stream": false,
	}
	
	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}
	
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqJSON))
	if err != nil {
		return "", err
	}
	
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	
	var respBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return "", err
	}
	
	choices, ok := respBody["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", nil
	}
	
	choice, ok := choices[0].(map[string]interface{})
	if !ok {
		return "", nil
	}
	
	message, ok := choice["message"].(map[string]interface{})
	if !ok {
		return "", nil
	}
	
	summary, ok := message["content"].(string)
	if !ok {
		return "", nil
	}
	
	return summary, nil
}
