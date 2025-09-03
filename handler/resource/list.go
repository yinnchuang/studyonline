package resource

import (
	"net/http"
	"strconv"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

func ListResource(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "3")
	pageStr := c.DefaultQuery("page", "0")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	offset := page * limit
	resourceWithLimit, err := service.ListResourceWithLimitOffset(c, limit, offset)
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

	limitStr := c.DefaultQuery("limit", "3")
	pageStr := c.DefaultQuery("page", "0")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	offset := page * limit

	// 展示所有
	if category == -1 {
		resourceWithLimit, err := service.ListResourceWithLimitOffset(c, limit, offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "请求成功",
				"data":    resourceWithLimit,
			})
		}
	}

	// 展示特定种类
	resourceWithCategory, err := service.ListResourceWithCategoryLimitOffset(c, limit, offset, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "请求成功",
			"data":    resourceWithCategory,
		})
	}
}

func ListResourceByUnit(c *gin.Context) {
	UnitStr := c.DefaultQuery("unit", "0")
	unit, _ := strconv.Atoi(UnitStr)

	limitStr := c.DefaultQuery("limit", "3")
	pageStr := c.DefaultQuery("page", "0")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	offset := page * limit

	// 展示特定种类
	resourceWithUnit, err := service.ListResourceWithUnitLimitOffset(c, limit, offset, unit)
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
