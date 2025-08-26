package main

import (
	"studyonline/constant"
	"studyonline/handler/admin"
	"studyonline/handler/dataset"
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
	// 资源
	v2 := r.Group("/resource")
	{
		v2.GET("/list", middleware.Auth(constant.CommonIdentity), resource.ListResource)
		v2.GET("/list/by/category", middleware.Auth(constant.CommonIdentity), resource.ListResource)
		// 一般是先upload然后获取地址，然后再create
		v2.POST("/upload", middleware.Auth(constant.TeacherIdentity), resource.UploadResource)
		v2.POST("/create", middleware.Auth(constant.TeacherIdentity), resource.CreateResource)
	}
	// 数据集
	v3 := r.Group("/dataset")
	{
		v3.GET("/list", dataset.ListDataset)

	}

	r.Run(":18000")
}
