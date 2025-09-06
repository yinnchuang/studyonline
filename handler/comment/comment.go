package comment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"studyonline/dao/entity"
	"studyonline/service"
)

func GetCommentByDiscussId(c *gin.Context) {
	discussIdStr := c.DefaultQuery("discussId", "0")
	discussId, _ := strconv.Atoi(discussIdStr)
	comments, err := service.GetCommentByDiscussId(c, uint(discussId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    comments,
	})
}

type CreateCommentDTO struct {
	DiscussId uint   `json:"discuss_id"`
	Comment   string `json:"comment"`
}

func CreateComment(c *gin.Context) {
	createCommentDTO := CreateCommentDTO{}
	err := c.ShouldBindBodyWithJSON(&createCommentDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	userId := c.GetUint("user_id")
	identity := c.GetInt("identity")
	userInfo, err := service.GetUserInfo(userId, identity)
	if err != nil || userInfo == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	comment := entity.Comment{
		UserId:    userId,
		DiscussId: createCommentDTO.DiscussId,
		Comment:   createCommentDTO.Comment,
		OwnerName: userInfo.Username,
	}
	err = service.CreateComment(c, comment)
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

type RemoveCommentDTO struct {
	CommentId uint `json:"comment_id"`
}

func RemoveComment(c *gin.Context) {
	removeCommentDTO := RemoveCommentDTO{}
	err := c.ShouldBindBodyWithJSON(&removeCommentDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	userId := c.GetUint("userId")
	err = service.RemoveComment(c, removeCommentDTO.CommentId, userId)
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
