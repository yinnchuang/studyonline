package resource

import (
	"net/http"
	"strconv"
	"studyonline/constant"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

type ListResourceVO struct {
	ID          uint        `json:"id"`
	Name        string      `json:"name"`
	CategoryID  int         `json:"category_id"`
	Description string      `json:"description,omitempty"`
	Units       interface{} `json:"unit_ids"`
}

func ListResource(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "-1")
	pageStr := c.DefaultQuery("page", "0")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	offset := page * limit
	resourceWithLimitOffset, err := service.ListResourceWithLimitOffset(c, limit, offset)

	// 不返回文件路径
	listResourceVOs := []ListResourceVO{}
	for _, item := range resourceWithLimitOffset {
		listResourceVOs = append(listResourceVOs, ListResourceVO{
			ID:          item.ID,
			Name:        item.Name,
			CategoryID:  item.CategoryID,
			Description: item.Description,
			Units:       item.Units,
		})
	}

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
		"data":    listResourceVOs,
		"total":   total,
	})
}

func ListResourceByCategory(c *gin.Context) {
	categoryStr := c.DefaultQuery("category", "-1")
	category, _ := strconv.Atoi(categoryStr)

	limitStr := c.DefaultQuery("limit", "-1")
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

	// 不返回文件路径
	listResourceVOs := []ListResourceVO{}
	for _, item := range resourceWithCategory {
		listResourceVOs = append(listResourceVOs, ListResourceVO{
			ID:          item.ID,
			Name:        item.Name,
			CategoryID:  item.CategoryID,
			Description: item.Description,
			Units:       item.Units,
		})
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    listResourceVOs,
		"total":   total,
	})

}

type ListResourceByUnitDTO struct {
	UnitIds []uint `json:"unit_ids"`
}

func ListResourceByUnit(c *gin.Context) {
	var listResourceByUnitDTO ListResourceByUnitDTO
	err := c.ShouldBindJSON(&listResourceByUnitDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	limitStr := c.DefaultQuery("limit", "-1")
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
