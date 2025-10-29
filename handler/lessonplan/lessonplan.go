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

func GenerateLessonPlan(c *gin.Context) {
	
}
