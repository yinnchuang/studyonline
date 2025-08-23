package admin

import (
	"net/http"
	"studyonline/constant"
	"studyonline/dao/entity"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

func ImportStudent(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	username := c.DefaultQuery("username", "")
	password := c.DefaultQuery("password", "")
	stu := entity.Student{
		Name:     name,
		Username: username,
		Password: password,
	}
	err := service.Import(c, stu, constant.StudentIdentity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}

func ImportTeacher(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	username := c.DefaultQuery("username", "")
	password := c.DefaultQuery("password", "")
	tea := entity.Teacher{
		Name:     name,
		Username: username,
		Password: password,
	}
	err := service.Import(c, tea, constant.TeacherIdentity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}
