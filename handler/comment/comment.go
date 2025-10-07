package comment

import (
	"net/http"
	"strconv"
	"studyonline/dao/entity"
	"studyonline/service"
	"time"

	"github.com/gin-gonic/gin"
)

type CommentVO struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	Name       string    `json:"name"`
	Username   string    `json:"user_name"`
	Department string    `json:"department"`
	DiscussId  uint      `json:"discuss_id"`
	Comment    string    `json:"comment"`
}

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
	var commentVOs []CommentVO
	for _, item := range comments {
		userInfo, err := service.GetUserInfo(item.UserId, item.Identity)
		if err != nil {
			continue
		}
		commentVOs = append(commentVOs, CommentVO{
			ID:         item.ID,
			CreatedAt:  item.CreatedAt,
			Name:       userInfo.Name,
			Username:   userInfo.Username,
			Department: userInfo.Department,
			DiscussId:  item.DiscussId,
			Comment:    item.Comment,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    commentVOs,
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
	userId := c.GetUint("userId")
	identity := c.GetInt("identity")

	comment := entity.Comment{
		UserId:    userId,
		DiscussId: createCommentDTO.DiscussId,
		Comment:   createCommentDTO.Comment,
		Identity:  identity,
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
	identity := c.GetInt("identity")
	comment, err := service.GetCommentById(c, removeCommentDTO.CommentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	if comment.UserId == userId && comment.Identity == identity {
		err = service.RemoveComment(c, removeCommentDTO.CommentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "请求成功",
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
}
