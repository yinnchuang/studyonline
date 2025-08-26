package dataset

import (
	"net/http"
	"strconv"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

func ListDataset(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "3")
	limit, _ := strconv.Atoi(limitStr)
	datasetWithLimit, err := service.ListDatasetWithLimit(c, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    datasetWithLimit,
	})
}
