package resource

import (
	"net/http"
	"strconv"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

func ListResource(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "3")
	limit, _ := strconv.Atoi(limitStr)
	resourceWithLimit, err := service.ListResourceWithLimit(c, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    resourceWithLimit,
	})
}

func ListResourceByCategory(c *gin.Context) {
	categoryStr := c.DefaultQuery("category", "-1")
	category, _ := strconv.Atoi(categoryStr)

	// 展示所有
	if category == -1 {

	}

	// 展示特定种类

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "请求失败",
	})
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}
