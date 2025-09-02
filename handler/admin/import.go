package admin

import (
	"net/http"
	"studyonline/constant"
	"studyonline/dao/entity"
	"studyonline/service"
	"studyonline/util"

	"github.com/gin-gonic/gin"
)

type ImportStudentPSO struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ImportStudent(c *gin.Context) {
	importStudentPSO := ImportStudentPSO{}
	err := c.ShouldBind(&importStudentPSO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	name := importStudentPSO.Name
	username := importStudentPSO.Username
	password := importStudentPSO.Password

	bcryptPassword, _ := util.GetPwd(password)
	stu := entity.Student{
		Name:     name,
		Username: username,
		Password: string(bcryptPassword),
	}
	err = service.Import(c, stu, constant.StudentIdentity)
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

type ImportTeacherPSO struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ImportTeacher(c *gin.Context) {
	importTeacherPSO := ImportTeacherPSO{}
	err := c.ShouldBind(&importTeacherPSO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
	}
	name := importTeacherPSO.Name
	username := importTeacherPSO.Username
	password := importTeacherPSO.Password

	bcryptPassword, _ := util.GetPwd(password)
	tea := entity.Teacher{
		Name:     name,
		Username: username,
		Password: string(bcryptPassword),
	}
	err = service.Import(c, tea, constant.TeacherIdentity)
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

func ListStudent(c *gin.Context) {
	students, err := service.List(c, constant.StudentIdentity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    students,
	})
}

func ListTeacher(c *gin.Context) {
	teachers, err := service.List(c, constant.TeacherIdentity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    teachers,
	})
}
