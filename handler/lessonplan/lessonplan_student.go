package lessonplan

import (
	"net/http"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

// 学生端看到的智能教案必须是老师发布后的，因此要单独区分

func GetAllLessonPlanStudent(c *gin.Context) {
	lessonPlans, err := service.GetAllLessonPlanStudent()
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
