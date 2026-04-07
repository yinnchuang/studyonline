package lessonplan_discuss

import (
	"net/http"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

type SummaryRequest struct {
	Content string `json:"content"`
}

func GetSummary(c *gin.Context) {
	summaryRequest := SummaryRequest{}
	err := c.ShouldBindBodyWithJSON(&summaryRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	
	summary, err := service.GetLessonPlanDiscussSummary(c, summaryRequest.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    summary,
	})
}
