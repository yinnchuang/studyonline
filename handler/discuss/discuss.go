package discuss

import (
	"net/http"
	"studyonline/dao/entity"
	"studyonline/service"
	"time"

	"github.com/gin-gonic/gin"
)

type DiscussesVO struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Username    string    `json:"user_name"`
	Department  string    `json:"department"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

func GetAllDiscusses(c *gin.Context) {
	discusses, err := service.GetAllDiscusses(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	var discussesVOs []DiscussesVO
	for _, item := range discusses {
		userInfo, err := service.GetUserInfo(item.UserId, item.Identity)
		if err != nil {
			continue
		}
		discussesVOs = append(discussesVOs, DiscussesVO{
			ID:          item.ID,
			CreatedAt:   item.CreatedAt,
			Name:        userInfo.Name,
			Username:    userInfo.Username,
			Department:  userInfo.Department,
			Title:       item.Title,
			Description: item.Description,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    discussesVOs,
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
		Identity:    identity,
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
	discuss, err := service.GetDiscuss(c, removeDiscussDTO.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	userId := c.GetUint("userId")
	identity := c.GetInt("identity")
	if discuss.UserId == userId && discuss.Identity == identity {
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
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
}
