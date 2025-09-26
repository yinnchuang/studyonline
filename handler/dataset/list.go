package dataset

import (
	"fmt"
	"net/http"
	"strconv"
	"studyonline/constant"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

type ListDatasetVO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	CategoryID  int    `json:"category_id"`
	Description string `json:"description,omitempty"`
	Scale       string `json:"scale"`
	Private     bool   `json:"private"`
}

func ListDataset(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "3")
	pageStr := c.DefaultQuery("page", "0")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	offset := page * limit
	datasetWithLimitOffset, err := service.ListDatasetWithLimitOffset(c, limit, offset)

	listDatasetVOs := []ListDatasetVO{}
	for _, item := range datasetWithLimitOffset {
		listDatasetVOs = append(listDatasetVOs, ListDatasetVO{
			ID:          item.ID,
			Name:        item.Name,
			CategoryID:  item.CategoryID,
			Description: item.Description,
			Scale:       item.Scale,
			Private:     item.Private,
		})
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	total, err := service.CountDataset(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    listDatasetVOs,
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
		return
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
	fmt.Println(datasetWithCategory)
	fmt.Println(total)
	listDatasetVOs := []ListDatasetVO{}
	for _, item := range datasetWithCategory {
		listDatasetVOs = append(listDatasetVOs, ListDatasetVO{
			ID:          item.ID,
			Name:        item.Name,
			CategoryID:  item.CategoryID,
			Description: item.Description,
			Scale:       item.Scale,
			Private:     item.Private,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    listDatasetVOs,
		"total":   total,
	})
}

func ListDatasetByTeacherId(c *gin.Context) {
	userId := c.GetUint("userId")

	// 根据teacherId查询数据集
	datasets, err := service.ListDatasetByTeacherId(c, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	listDatasetVOs := []ListDatasetVO{}
	for _, item := range datasets {
		listDatasetVOs = append(listDatasetVOs, ListDatasetVO{
			ID:          item.ID,
			Name:        item.Name,
			CategoryID:  item.CategoryID,
			Description: item.Description,
			Scale:       item.Scale,
			Private:     item.Private,
		})
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    listDatasetVOs,
	})
}
