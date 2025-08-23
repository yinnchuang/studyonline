package main

import (
	"net/http"
	"studyonline/handler/login"
	"studyonline/handler/middleware"
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
		v0.GET("/admin", login.AdminLogin)
	}
	// 获取资源
	v1 := r.Group("/resource", middleware.Auth)
	{
		v1.GET("/file", middleware.Auth, func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})
		v1.GET("/dataset", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})
	}
	r.Run(":8000")
}
