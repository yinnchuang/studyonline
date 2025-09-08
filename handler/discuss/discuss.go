package discuss

import (
	"net/http"
	"studyonline/dao/entity"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

func GetAllDiscusses(c *gin.Context) {
	discusses, err := service.GetAllDiscusses(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    discusses,
	})
}

type CreateDiscussDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateDiscuss(c *gin.Context) {
	createDiscussDTO := CreateDiscussDTO{}
	err := c.ShouldBindBodyWithJSON(&createDiscussDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	userId := c.GetUint("userId")
	identity := c.GetInt("identity")
	userInfo, err := service.GetUserInfo(userId, identity)
	if err != nil || userInfo == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	discuss := entity.Discuss{
		UserId:      userId,
		Title:       createDiscussDTO.Title,
		Description: createDiscussDTO.Description,
		OwnerName:   userInfo.Username,
	}
	err = service.CreateDiscuss(c, discuss)
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

type RemoveDiscussDTO struct {
	Id uint `json:"id"`
}

func RemoveDiscuss(c *gin.Context) {
	removeDiscussDTO := RemoveDiscussDTO{}
	err := c.ShouldBindBodyWithJSON(&removeDiscussDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	err = service.RemoveDiscuss(c, removeDiscussDTO.Id)
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
