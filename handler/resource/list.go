package resource

import (
	"fmt"
	"net/http"
	"strconv"
	"studyonline/constant"
	"studyonline/service"
	"time"

	"github.com/gin-gonic/gin"
)

type ListResourceVO struct {
	ID           uint        `json:"id"`
	CreatedAt    time.Time   `json:"created_at"`
	Name         string      `json:"name"`
	CategoryID   int         `json:"category_id"`
	Description  string      `json:"description,omitempty"`
	Units        interface{} `json:"unit_ids"`
	DownloadTime int         `json:"download_time"`
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
			ID:           item.ID,
			CreatedAt:    item.CreatedAt,
			Name:         item.Name,
			CategoryID:   item.CategoryID,
			Description:  item.Description,
			Units:        item.Units,
			DownloadTime: item.DownloadTime,
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
		return
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
			ID:           item.ID,
			CreatedAt:    item.CreatedAt,
			Name:         item.Name,
			CategoryID:   item.CategoryID,
			Description:  item.Description,
			Units:        item.Units,
			DownloadTime: item.DownloadTime,
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
	UnitIds []uint `form:"unit_ids"`
}

func ListResourceByUnit(c *gin.Context) {
	var listResourceByUnitDTO ListResourceByUnitDTO
	err := c.ShouldBindQuery(&listResourceByUnitDTO)
	fmt.Println(listResourceByUnitDTO)
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

	listResourceVOs := []ListResourceVO{}
	for _, item := range resourceWithUnit {
		listResourceVOs = append(listResourceVOs, ListResourceVO{
			ID:           item.ID,
			CreatedAt:    item.CreatedAt,
			Name:         item.Name,
			CategoryID:   item.CategoryID,
			Description:  item.Description,
			Units:        item.Units,
			DownloadTime: item.DownloadTime,
		})
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "请求成功",
			"data":    listResourceVOs,
		})
	}
}

// SearchResourceByKeyword 仅通过关键词搜索资源（匹配名称或描述）
func SearchResourceByKeyword(c *gin.Context) {
	// 解析分页参数（保持与原有接口一致）
	limitStr := c.DefaultQuery("limit", "-1")
	pageStr := c.DefaultQuery("page", "0")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	offset := page * limit

	// 获取搜索关键词（从query参数中获取）
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "搜索关键词不能为空",
		})
		return
	}

	// 调用服务层根据关键词搜索（匹配名称或描述）
	resources, err := service.SearchResourceByKeyword(c, limit, offset, keyword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	// 转换为VO（隐藏不需要的字段，复用原有结构）
	listResourceVOs := []ListResourceVO{}
	for _, item := range resources {
		listResourceVOs = append(listResourceVOs, ListResourceVO{
			ID:           item.ID,
			CreatedAt:    item.CreatedAt,
			Name:         item.Name,
			CategoryID:   item.CategoryID,
			Description:  item.Description,
			Units:        item.Units,
			DownloadTime: item.DownloadTime,
		})
	}

	// 获取符合条件的总条数
	total, err := service.CountResourceByKeyword(c, keyword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "搜索成功",
		"data":    listResourceVOs,
		"total":   total,
	})
}
