package unit

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"studyonline/dao/entity"
	"studyonline/service"
)

func GetAllUnit(c *gin.Context) {
	units, err := service.GetAllUnit(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"units":   units,
	})
}

func RemoveUnit(c *gin.Context) {
	unitIdStr := c.DefaultQuery("id", "")
	if unitIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	unitId, err := strconv.Atoi(unitIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	err = service.RemoveUnit(c, uint(unitId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}

type CreateUnitPSO struct {
	UnitName string `json:"unit_name" binding:"required"`
	UnitDesc string `json:"unit_desc"`
}

func CreateUnit(c *gin.Context) {
	var createUnitPSO CreateUnitPSO
	if err := c.ShouldBindJSON(&createUnitPSO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	unit := entity.Unit{
		UnitName: createUnitPSO.UnitName,
		UnitDesc: createUnitPSO.UnitDesc,
	}
	err := service.CreateUnit(c, unit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}
