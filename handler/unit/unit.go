package unit

import (
	"net/http"
	"strconv"
	"studyonline/dao/entity"
	"studyonline/service"

	"github.com/gin-gonic/gin"
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

type CreateUnitDTO struct {
	UnitName     string `json:"unit_name" binding:"required"`
	UnitDesc     string `json:"unit_desc"`
	FatherUnitId uint   `json:"father_unit_id" binding:"required"`
}

func CreateUnit(c *gin.Context) {
	var createUnitDTO CreateUnitDTO
	if err := c.ShouldBindJSON(&createUnitDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	unit := entity.Unit{
		UnitName:     createUnitDTO.UnitName,
		UnitDesc:     createUnitDTO.UnitDesc,
		FatherUnitId: createUnitDTO.FatherUnitId,
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
