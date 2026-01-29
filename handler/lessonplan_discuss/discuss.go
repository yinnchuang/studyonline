package lessonplan_discuss

import (
	"net/http"
	"strconv"
	"studyonline/dao/entity"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

func GetDiscussesByLessonPlanID(c *gin.Context) {
	lessonPlanIDStr := c.Query("LessonPlanID")
	lessonPlanID, _ := strconv.Atoi(lessonPlanIDStr)
	discusses, err := service.GetLessonPlanDiscussByLessonPlanID(c, uint(lessonPlanID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    discusses,
	})
}

type CreateLessonPlanDiscussDTO struct {
	LessonPlanID uint   `json:"lesson_plan_id"` // 绑定某个知识点
	FatherID     uint   `json:"father_id"`      // 父讨论id
	Content      string `json:"content"`
}

func CreateLessonPlanDiscuss(c *gin.Context) {
	createLessonPlanDiscussDTO := CreateLessonPlanDiscussDTO{}
	err := c.ShouldBindBodyWithJSON(&createLessonPlanDiscussDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	userId := c.GetUint("userId")
	identity := c.GetInt("identity")
	userInfo, err := service.GetUserInfo(userId, identity)
	if err != nil || userInfo == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	discuss := entity.LessonPlanDiscuss{
		UserId:       userId,
		Identity:     identity,
		LessonPlanID: createLessonPlanDiscussDTO.LessonPlanID,
		Content:      createLessonPlanDiscussDTO.Content,
		FatherID:     createLessonPlanDiscussDTO.FatherID,
	}
	err = service.CreateLessonPlanDiscuss(c, discuss)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}

type RemoveLessonPlanDiscussDTO struct {
	Id uint `json:"id"`
}

func RemoveLessonPlanDiscuss(c *gin.Context) {
	removeLessonPlanDiscussDTO := RemoveLessonPlanDiscussDTO{}
	err := c.ShouldBindBodyWithJSON(&removeLessonPlanDiscussDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	discuss, err := service.GetLessonPlanDiscussByID(c, removeLessonPlanDiscussDTO.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	userId := c.GetUint("userId")
	identity := c.GetInt("identity")
	if discuss.UserId == userId && discuss.Identity == identity {
		err = service.RemoveLessonPlanDiscussByID(c, removeLessonPlanDiscussDTO.Id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
		err = service.RemoveLessonPlanDiscussByFatherID(c, discuss.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "请求成功",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
}
