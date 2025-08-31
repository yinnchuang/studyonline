package main

import (
	"studyonline/constant"
	"studyonline/handler/admin"
	"studyonline/handler/announcement"
	"studyonline/handler/dataset"
	"studyonline/handler/login"
	"studyonline/handler/middleware"
	"studyonline/handler/resource"
	"studyonline/handler/score"
	"studyonline/handler/unit"
	minit "studyonline/init"

	"github.com/gin-gonic/gin"
)

func init() {
	minit.Init()
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
		v1.GET("/list/student", middleware.Auth(constant.AdminIdentity), admin.ListStudent)
		v1.GET("/list/teacher", middleware.Auth(constant.AdminIdentity), admin.ListTeacher)
		v1.POST("/import/student", middleware.Auth(constant.AdminIdentity), admin.ImportStudent)
		v1.POST("/import/teacher", middleware.Auth(constant.AdminIdentity), admin.ImportTeacher)
	}
	// 资源
	v2 := r.Group("/resource")
	{
		v2.GET("/list", middleware.Auth(constant.CommonIdentity), resource.ListResource)
		v2.GET("/list/by/category", middleware.Auth(constant.CommonIdentity), resource.ListResourceByCategory)
		// 一般是先upload文件然后获取地址，然后再create
		v2.POST("/upload", middleware.Auth(constant.TeacherIdentity), resource.UploadResource)
		v2.POST("/create", middleware.Auth(constant.TeacherIdentity), resource.CreateResource)
	}
	// 数据集
	v3 := r.Group("/dataset")
	{
		v3.GET("/list", middleware.Auth(constant.CommonIdentity), dataset.ListDataset)
		v3.GET("/list/by/category", middleware.Auth(constant.CommonIdentity), dataset.ListDatasetByCategory)
		// 一般是先upload文件然后获取地址，然后再create
		v3.POST("/upload", middleware.Auth(constant.TeacherIdentity), dataset.UploadDataset)
		v3.POST("/create", middleware.Auth(constant.TeacherIdentity), dataset.CreateDataset)
	}
	// 知识点
	v4 := r.Group("/unit")
	{
		v4.GET("/list", middleware.Auth(constant.CommonIdentity), unit.GetAllUnit)
		v4.POST("/create", middleware.Auth(constant.TeacherIdentity), unit.CreateUnit)
		v4.POST("/delete", middleware.Auth(constant.TeacherIdentity), unit.RemoveUnit)
	}
	// 公告
	v5 := r.Group("/announcement")
	{
		v5.GET("/list", middleware.Auth(constant.CommonIdentity), announcement.GetAllAnnouncements)
		v5.POST("/create", middleware.Auth(constant.TeacherIdentity), announcement.CreateAnnouncement)
		v5.POST("/remove", middleware.Auth(constant.TeacherIdentity), announcement.RemoveAnnouncement)
	}
	// 分数
	v6 := r.Group("/score")
	{
		v6.GET("/list", middleware.Auth(constant.TeacherIdentity), score.GetAllScore)
		v6.GET("/list/mean", middleware.Auth(constant.TeacherIdentity), score.GetMeanScore)
		v6.GET("/list/student", middleware.Auth(constant.StudentIdentity), score.GetScoreByStudentId)
	}
	r.Run(":18000")
}
