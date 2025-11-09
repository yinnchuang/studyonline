package user

import (
	"fmt"
	"net/http"
	"studyonline/constant"
	"studyonline/dao/redis"
	"studyonline/service"
	"studyonline/util"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	studentId := c.GetUint("userId")
	identity := c.GetInt("identity")
	if identity == constant.StudentIdentity || identity == constant.TeacherIdentity {
		userInfo, err := service.GetUserInfo(studentId, identity)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "请求成功",
				"data":    userInfo,
			})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "请求失败",
	})
	return
}

type ChangePasswordDTO struct {
	Password string `json:"password"`
}

func ChangePassword(c *gin.Context) {
	var changePasswordDTO ChangePasswordDTO
	if err := c.ShouldBindJSON(&changePasswordDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	userId := c.GetUint("userId")
	identity := c.GetInt("identity")
	bcryptPassword, _ := util.GetPwd(changePasswordDTO.Password)

	if !util.IsValidPassword(changePasswordDTO.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "密码过于简单",
		})
		return
	}

	if identity == constant.StudentIdentity {
		err := service.ChangeStudentPassword(userId, string(bcryptPassword))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
		userInfo, err := service.GetUserInfo(userId, identity)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
		cacheKey := fmt.Sprintf("change_password_%v_%v", userInfo.Username, constant.StudentIdentity)
		redis.RDB.Set(c, cacheKey, 1, time.Hour*24*7)
	} else if identity == constant.TeacherIdentity {
		err := service.ChangeTeacherPassword(userId, string(bcryptPassword))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
		userInfo, err := service.GetUserInfo(userId, identity)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
		cacheKey := fmt.Sprintf("change_password_%v_%v", userInfo.Username, constant.TeacherIdentity)
		redis.RDB.Set(c, cacheKey, 1, time.Hour*24*7)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}

type BindEmailDTO struct {
	Email string `json:"email"`
}

func BindEmail(c *gin.Context) {
	var bindEmailDTO BindEmailDTO
	if err := c.ShouldBindJSON(&bindEmailDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	if !util.IsValidEmail(bindEmailDTO.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "非邮箱格式",
		})
		return
	}
	userId := c.GetUint("userId")
	identity := c.GetInt("identity")

	if identity == constant.StudentIdentity {
		err := service.BindStudentEmail(userId, bindEmailDTO.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
	} else if identity == constant.TeacherIdentity {
		err := service.BindTeacherEmail(userId, bindEmailDTO.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}

type SendCodeDTO struct {
	Username string `json:"username"`
}

func SendCodeStudent(c *gin.Context) {
	var sendCodeDTO SendCodeDTO
	if err := c.ShouldBindJSON(&sendCodeDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	student, err := service.GetStudentInfoByUsername(sendCodeDTO.Username)
	if err != nil || student == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不存在",
		})
		return
	}
	if !util.IsValidEmail(student.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未绑定邮箱或邮箱无效",
		})
		return
	}
	err = service.SendCode2Email(c, student.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "发送失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}

func SendCodeTeacher(c *gin.Context) {
	var sendCodeDTO SendCodeDTO
	if err := c.ShouldBindJSON(&sendCodeDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	teacher, err := service.GetTeacherInfoByUsername(sendCodeDTO.Username)
	if err != nil || teacher == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不存在",
		})
		return
	}
	if !util.IsValidEmail(teacher.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未绑定邮箱或邮箱无效",
		})
		return
	}
	err = service.SendCode2Email(c, teacher.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "发送失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}

type ChangePasswordByEmailDTO struct {
	Username string `json:"username"`
	Code     string `json:"code"`
	Password string `json:"password"`
}

func ChangePasswordByEmailStudent(c *gin.Context) {
	var changePasswordByEmailDTO ChangePasswordByEmailDTO
	if err := c.ShouldBindJSON(&changePasswordByEmailDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	if !util.IsValidPassword(changePasswordByEmailDTO.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "密码过于简单",
		})
		return
	}
	student, err := service.GetStudentInfoByUsername(changePasswordByEmailDTO.Username)
	if err != nil || student == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不存在",
		})
		return
	}
	code := redis.RDB.Get(c, student.Email).String()
	if code != changePasswordByEmailDTO.Code {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "验证码错误",
		})
		return
	}
	bcryptPassword, _ := util.GetPwd(changePasswordByEmailDTO.Password)
	err = service.ChangeStudentPassword(student.ID, string(bcryptPassword))
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

func ChangePasswordByEmailTeacher(c *gin.Context) {
	var changePasswordByEmailDTO ChangePasswordByEmailDTO
	if err := c.ShouldBindJSON(&changePasswordByEmailDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
		})
		return
	}
	if !util.IsValidPassword(changePasswordByEmailDTO.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "密码过于简单",
		})
		return
	}
	teacher, err := service.GetTeacherInfoByUsername(changePasswordByEmailDTO.Username)
	if err != nil || teacher == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "用户不存在",
		})
		return
	}
	code := redis.RDB.Get(c, teacher.Email).String()
	if code != changePasswordByEmailDTO.Code {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "验证码错误",
		})
		return
	}
	bcryptPassword, _ := util.GetPwd(changePasswordByEmailDTO.Password)
	err = service.ChangeTeacherPassword(teacher.ID, string(bcryptPassword))
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
