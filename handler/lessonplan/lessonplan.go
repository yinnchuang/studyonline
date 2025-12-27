package lessonplan

import (
	"net/http"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

func GetAllLessonPlan(c *gin.Context) {
	lessonPlans, err := service.GetAllLessonPlan()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    lessonPlans,
	})
}

type RemoveLessonPlanDTO struct {
	LessonPlanId uint `json:"lesson_plan_id"`
}

func RemoveLessonPlan(c *gin.Context) {
	removeLessonPlanDTO := RemoveLessonPlanDTO{}
	err := c.ShouldBindBodyWithJSON(&removeLessonPlanDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	err = service.RemoveLessonPlan(c, removeLessonPlanDTO.LessonPlanId)
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
