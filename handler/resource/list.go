package resource

import (
	"net/http"
	"strconv"
	"studyonline/constant"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

func ListResource(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "3")
	pageStr := c.DefaultQuery("page", "0")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	offset := page * limit
	resourceWithLimitOffset, err := service.ListResourceWithLimitOffset(c, limit, offset)
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
		"data":    resourceWithLimitOffset,
		"total":   total,
	})
}

func ListResourceByCategory(c *gin.Context) {
	categoryStr := c.DefaultQuery("category", "-1")
	category, _ := strconv.Atoi(categoryStr)

	limitStr := c.DefaultQuery("limit", "3")
	pageStr := c.DefaultQuery("page", "0")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	offset := page * limit

	// 非类别之一
	if !constant.IfResourceCategory(category) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
	}

	// 展示特定种类
	resourceWithCategory, err := service.ListResourceWithCategoryLimitOffset(c, limit, offset, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	total, err := service.CountResourceWithCategory(c, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    resourceWithCategory,
		"total":   total,
	})

}

type ListResourceByUnitDTO struct {
	UnitIds []uint `json:"unit_ids"`
}

func ListResourceByUnit(c *gin.Context) {
	var listResourceByUnitDTO ListResourceByUnitDTO
	err := c.ShouldBind(&listResourceByUnitDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	limitStr := c.DefaultQuery("limit", "3")
	pageStr := c.DefaultQuery("page", "0")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	offset := page * limit

	// 展示特定种类
	resourceWithUnit, err := service.ListResourceWithUnitLimitOffset(c, limit, offset, listResourceByUnitDTO.UnitIds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "请求成功",
			"data":    resourceWithUnit,
		})
	}
}
