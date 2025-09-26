package admin

import (
	"net/http"
	"studyonline/constant"
	"studyonline/dao/entity"
	"studyonline/service"
	"studyonline/util"

	"github.com/gin-gonic/gin"
)

type ImportStudentDTO struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ImportStudent(c *gin.Context) {
	importStudentDTO := ImportStudentDTO{}
	err := c.ShouldBind(&importStudentDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	name := importStudentDTO.Name
	username := importStudentDTO.Username
	password := importStudentDTO.Password

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

type ImportTeacherDTO struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ImportTeacher(c *gin.Context) {
	importTeacherDTO := ImportTeacherDTO{}
	err := c.ShouldBind(&importTeacherDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
	}
	name := importTeacherDTO.Name
	username := importTeacherDTO.Username
	password := importTeacherDTO.Password

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
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    teachers,
	})
}

type DeleteStudentDTO struct {
	StudentId uint `json:"student_id"`
}

func DeleteStudent(c *gin.Context) {
	var deleteStudentDTO DeleteStudentDTO
	err := c.ShouldBind(&deleteStudentDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	err = service.DeleteStudent(c, deleteStudentDTO.StudentId)
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

type DeleteTeacherDTO struct {
	TeacherId uint `json:"teacher_id"`
}

func DeleteTeacher(c *gin.Context) {
	var deleteTeacherDTO DeleteTeacherDTO
	err := c.ShouldBind(&deleteTeacherDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	err = service.DeleteTeacher(c, deleteTeacherDTO.TeacherId)
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
