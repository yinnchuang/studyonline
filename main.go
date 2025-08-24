package main

import (
	"net/http"
	"studyonline/constant"
	"studyonline/handler/admin"
	"studyonline/handler/login"
	"studyonline/handler/middleware"
	"studyonline/handler/resource"
	"studyonline/init"

	"github.com/gin-gonic/gin"
)

func init() {
	init.Init()
}

func main() {
	r := gin.Default()
	// 登录
	v0 := r.Group("/login")
	{
		v0.POST("/student", login.StudentLogin)
		v0.POST("/teacher", login.TeacherLogin)
		v0.POST("/admin", login.AdminLogin)
	}
	// 管理员导入
	v1 := r.Group("/admin")
	{
		v1.POST("/import/student", middleware.Auth(constant.AdminIdentity), admin.ImportStudent)
		v1.POST("/import/teacher", middleware.Auth(constant.AdminIdentity), admin.ImportTeacher)
	}
	// 获取资源
	v2 := r.Group("/resource")
	{
		v2.GET("/list/random", resource.RandomListResource)
		
		v2.GET("/file", middleware.Auth(constant.CommonIdentity), func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})
		v2.GET("/dataset", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})
	}
	// 上传资源

	r.Run(":18000")
}
