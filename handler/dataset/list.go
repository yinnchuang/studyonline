package dataset

import (
	"net/http"
	"strconv"
	"studyonline/constant"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

func ListDataset(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "3")
	pageStr := c.DefaultQuery("page", "0")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	offset := page * limit
	datasetWithLimit, err := service.ListDatasetWithLimitOffset(c, limit, offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	total, err := service.CountResource(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    datasetWithLimit,
		"total":   total,
	})
}

func ListDatasetByCategory(c *gin.Context) {
	categoryStr := c.DefaultQuery("category", "-1")
	category, _ := strconv.Atoi(categoryStr)

	limitStr := c.DefaultQuery("limit", "3")
	pageStr := c.DefaultQuery("page", "0")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	offset := page * limit

	// 非类别之一
	if !constant.IfDatasetCategory(category) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
	}

	// 展示特定种类
	datasetWithCategory, err := service.ListDatasetWithCategoryLimitOffset(c, limit, offset, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	total, err := service.CountDatasetWithCategory(c, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    datasetWithCategory,
		"total":   total,
	})
}

func ListDatasetByUnit(c *gin.Context) {
	unitStr := c.DefaultQuery("unit", "-1")
	unit, _ := strconv.Atoi(unitStr)

	// 展示特定种类
	datasetWithUnit, err := service.ListDatasetWithUnit(c, unit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "请求成功",
			"data":    datasetWithUnit,
		})
	}
}
