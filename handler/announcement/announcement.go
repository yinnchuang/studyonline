package announcement

import (
	"net/http"
	"studyonline/dao/entity"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

func GetAllAnnouncements(c *gin.Context) {
	announcements, err := service.GetAllAnnouncement(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    announcements,
	})
}

type CreateAnnouncementDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateAnnouncement(c *gin.Context) {
	createAnnouncementDTO := CreateAnnouncementDTO{}
	err := c.ShouldBindBodyWithJSON(&createAnnouncementDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	announcement := entity.Announcement{
		Title:       createAnnouncementDTO.Title,
		Description: createAnnouncementDTO.Description,
	}
	err = service.CreateAnnouncement(c, announcement)
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

type RemoveAnnouncementDTO struct {
	Id uint `json:"id"`
}

func RemoveAnnouncement(c *gin.Context) {
	removeAnnouncementDTO := RemoveAnnouncementDTO{}
	err := c.ShouldBindBodyWithJSON(&removeAnnouncementDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	err = service.RemoveAnnouncement(c, removeAnnouncementDTO.Id)
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
