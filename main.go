package main

import (
	"studyonline/constant"
	"studyonline/handler/admin"
	"studyonline/handler/comment"
	"studyonline/handler/dataset"
	"studyonline/handler/discuss"
	"studyonline/handler/homework"
	"studyonline/handler/login"
	"studyonline/handler/middleware"
	"studyonline/handler/permission"
	"studyonline/handler/resource"
	"studyonline/handler/score"
	"studyonline/handler/unit"
	"studyonline/handler/user"
	minit "studyonline/init"
	"studyonline/log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	minit.Init()
}

func main() {
	defer log.DownloadLogger.Close() // 关闭log日志
	defer log.CommonLogger.Close()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 也可写具体前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 关键：加上你用的头
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // 如果带 cookie
		MaxAge:           12 * time.Hour,
	}))
	r.Use(middleware.RateLimitGlobal())

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
		v1.GET("/list/student", middleware.Auth(constant.CommonIdentity), admin.ListStudent)
		v1.GET("/list/teacher", middleware.Auth(constant.AdminIdentity), admin.ListTeacher)
		v1.POST("/import/student", middleware.Auth(constant.AdminIdentity), admin.ImportStudent)
		v1.POST("/import/teacher", middleware.Auth(constant.AdminIdentity), admin.ImportTeacher)
		v1.POST("/delete/student", middleware.Auth(constant.AdminIdentity), admin.DeleteStudent)
		v1.POST("/delete/teacher", middleware.Auth(constant.AdminIdentity), admin.DeleteTeacher)
	}
	// 资源
	v2 := r.Group("/resource")
	{
		v2.GET("/list", middleware.Auth(constant.CommonIdentity), resource.ListResource)
		v2.GET("/list/by/category", middleware.Auth(constant.CommonIdentity), resource.ListResourceByCategory)
		v2.GET("/list/by/unit", middleware.Auth(constant.CommonIdentity), resource.ListResourceByUnit)

		v2.POST("/uploadAndCreate", middleware.Auth(constant.TeacherIdentity), resource.UploadAndCreateResource)
		v2.GET("/cover", middleware.Auth(constant.CommonIdentity), resource.GetResourceCover)
		v2.GET("/file", middleware.Auth(constant.CommonIdentity), resource.GetResource)
	}
	// 数据集
	v3 := r.Group("/dataset")
	{
		v3.GET("/list", middleware.Auth(constant.CommonIdentity), dataset.ListDataset)
		v3.GET("/list/by/category", middleware.Auth(constant.CommonIdentity), dataset.ListDatasetByCategory)
		v3.GET("/list/by/teacherId", middleware.Auth(constant.TeacherIdentity), dataset.ListDatasetByTeacherId)

		v3.POST("/uploadAndCreate", middleware.Auth(constant.TeacherIdentity), dataset.UploadAndCreateDataset)
		v3.GET("/cover", middleware.Auth(constant.CommonIdentity), dataset.GetDatasetCover)
		v3.GET("/file", middleware.Auth(constant.CommonIdentity), dataset.GetDataset)
	}
	// 知识点
	v4 := r.Group("/unit")
	{
		v4.GET("/list", middleware.Auth(constant.CommonIdentity), unit.GetAllUnit)
		v4.POST("/create", middleware.Auth(constant.TeacherIdentity), unit.CreateUnit)
		v4.POST("/delete", middleware.Auth(constant.TeacherIdentity), unit.RemoveUnit)
	}
	// 申请权限
	v5 := r.Group("/permission")
	{
		// 同意请求
		v5.POST("/agree", middleware.Auth(constant.TeacherIdentity), permission.AgreePermission)
		// 拒绝请求
		v5.POST("/disagree", middleware.Auth(constant.TeacherIdentity), permission.DisagreePermission)
		// 发起请求
		v5.POST("/request", middleware.Auth(constant.CommonIdentity), permission.RequestPermissionByDatasetId)
		// 申请记录
		v5.GET("/list/request", middleware.Auth(constant.CommonIdentity), permission.ListRequestByUserId)
		// 审批记录
		v5.GET("/list/review", middleware.Auth(constant.TeacherIdentity), permission.ListRequestByTeacherId)
	}
	//// 公告
	//v5 := r.Group("/announcement")
	//{
	//	v5.GET("/list", middleware.Auth(constant.CommonIdentity), announcement.GetAllAnnouncements)
	//	v5.POST("/create", middleware.Auth(constant.TeacherIdentity), announcement.CreateAnnouncement)
	//	v5.POST("/remove", middleware.Auth(constant.TeacherIdentity), announcement.RemoveAnnouncement)
	//}
	// 作业
	v6 := r.Group("/homework")
	{
		v6.GET("/list", middleware.Auth(constant.CommonIdentity), homework.ListHomework)
		v6.POST("/uploadAndCreate", middleware.Auth(constant.TeacherIdentity), homework.UploadAndCreateHomework)
		v6.POST("/delete", middleware.Auth(constant.TeacherIdentity), homework.RemoveHomework)
		v6.GET("/file", middleware.Auth(constant.CommonIdentity), homework.GetHomework)
	}
	//// 学生的提交
	//v7 := r.Group("/submission")
	//{
	//	v7.GET("/list/by/homeworkId", middleware.Auth(constant.CommonIdentity), submission.ListSubmissionByHomeworkId)
	//	v7.POST("/upload", middleware.Auth(constant.StudentIdentity), submission.UploadSubmission)
	//	v7.POST("/create", middleware.Auth(constant.StudentIdentity), submission.CreateSubmission)
	//	v7.POST("/delete", middleware.Auth(constant.StudentIdentity), submission.RemoveSubmission)
	//}
	// 分数
	v8 := r.Group("/score")
	{
		// 给教师调用
		v8.GET("/list", middleware.Auth(constant.TeacherIdentity), score.GetAllScore)
		v8.POST("/create", middleware.Auth(constant.TeacherIdentity), score.CreateScore)
		// 根据studentId展示，给student调用
		v8.GET("/list/by/studentId", middleware.Auth(constant.StudentIdentity), score.GetScoreByStudentId)
	}
	// 讨论
	v9 := r.Group("/discuss")
	{
		v9.GET("/list", middleware.Auth(constant.CommonIdentity), discuss.GetAllDiscusses)
		v9.POST("/create", middleware.Auth(constant.CommonIdentity), discuss.CreateDiscuss)
		v9.POST("/delete", middleware.Auth(constant.CommonIdentity), discuss.RemoveDiscuss)
	}
	// 评论
	v10 := r.Group("/comment")
	{
		v10.GET("/list/by/discussId", middleware.Auth(constant.CommonIdentity), comment.GetCommentByDiscussId)
		v10.POST("/create", middleware.Auth(constant.CommonIdentity), comment.CreateComment)
		v10.POST("/delete", middleware.Auth(constant.CommonIdentity), comment.RemoveComment)
	}
	// 用户信息
	v11 := r.Group("/user")
	{
		v11.GET("/info", middleware.Auth(constant.CommonIdentity), user.GetUserInfo)
		v11.POST("/changePassword", middleware.Auth(constant.CommonIdentity), user.ChangePassword)
	}
	// 静态资源
	// r.Static("/static", "./static") // 废弃

	r.Run(":8080")
}
