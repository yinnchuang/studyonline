package admin

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"studyonline/constant"
	"studyonline/dao/entity"
	"studyonline/dao/redis"
	"studyonline/service"
	"studyonline/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type ImportStudentDTO struct {
	Name       string `json:"name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Department string `json:"department"`
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
	department := importStudentDTO.Department

	bcryptPassword, _ := util.GetPwd(password)
	stu := entity.Student{
		Name:       name,
		Username:   username,
		Password:   string(bcryptPassword),
		Department: department,
	}
	err = service.Import(c, stu, constant.StudentIdentity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	cacheKey := fmt.Sprintf("change_password_%v_%v", username, constant.StudentIdentity)
	redis.RDB.Set(c, cacheKey, -1, time.Hour*24*60)
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}

type ImportTeacherDTO struct {
	Name       string `json:"name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Department string `json:"department"`
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
	department := importTeacherDTO.Department

	bcryptPassword, _ := util.GetPwd(password)
	tea := entity.Teacher{
		Name:       name,
		Username:   username,
		Password:   string(bcryptPassword),
		Department: department,
	}
	err = service.Import(c, tea, constant.TeacherIdentity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	cacheKey := fmt.Sprintf("change_password_%v_%v", username, constant.TeacherIdentity)
	redis.RDB.Set(c, cacheKey, -1, time.Hour*24*60)
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}

func ListStudent(c *gin.Context) {
	fmt.Println("ListStudent")
	students, err := service.List(c, constant.StudentIdentity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
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

type ResetTeacherDTO struct {
	TeacherId uint   `json:"teacher_id"`
	Password  string `json:"password"`
}

func ResetTeacher(c *gin.Context) {
	var resetTeacherDTO ResetTeacherDTO
	err := c.ShouldBind(&resetTeacherDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	info, err := service.GetTeacherInfo(resetTeacherDTO.TeacherId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	if info == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	bcryptPassword, _ := util.GetPwd(resetTeacherDTO.Password)
	err = service.ChangeTeacherPassword(resetTeacherDTO.TeacherId, string(bcryptPassword))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	cacheKey := fmt.Sprintf("change_password_%v_%v", info.Username, constant.TeacherIdentity)
	redis.RDB.Set(c, cacheKey, -1, time.Hour*24*60)
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}

type ResetStudentDTO struct {
	StudentId uint   `json:"student_id"`
	Password  string `json:"password"`
}

func ResetStudent(c *gin.Context) {
	var resetStudentDTO ResetStudentDTO
	err := c.ShouldBind(&resetStudentDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	info, err := service.GetStudentInfo(resetStudentDTO.StudentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	if info == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	bcryptPassword, _ := util.GetPwd(resetStudentDTO.Password)
	err = service.ChangeStudentPassword(resetStudentDTO.StudentId, string(bcryptPassword))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	cacheKey := fmt.Sprintf("change_password_%v_%v", info.Username, constant.StudentIdentity)
	redis.RDB.Set(c, cacheKey, -1, time.Hour*24*60)
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}

func GetFormatExcel(c *gin.Context) {
	c.File("./static/format.xlsx")
}

func ImportStudentByExcel(c *gin.Context) {
	// 获取上传的文件
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	// 验证文件后缀（仅允许 .xlsx 格式，忽略大小写）
	filename := fileHeader.Filename
	// 将文件名转为小写，方便统一判断（如 .XLSX 会被转为 .xlsx）
	lowerFilename := strings.ToLower(filename)
	// 检查是否以 .xlsx 结尾
	if !strings.HasSuffix(lowerFilename, ".xlsx") {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "仅支持 .xlsx 格式的Excel文件",
		})
		return
	}

	// 打开文件流
	srcFile, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	defer srcFile.Close()

	// 解析Excel
	file, err := excelize.OpenReader(srcFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "解析Excel失败（文件可能损坏或格式错误）",
		})
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("关闭Excel文件失败: %v\n", err)
		}
	}()

	// 读取并打印数据
	sheetName := file.GetSheetList()[0]
	rows, err := file.GetRows(sheetName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	var students []entity.Student

	// 逐行遍历import
	for i, row := range rows {
		// 跳过标题行
		if i == 0 {
			continue
		}
		if len(row) < 4 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
		name := row[0]
		username := row[1]
		password := row[2]
		department := row[3]
		bcryptPassword, _ := util.GetPwd(password)
		stu := entity.Student{
			Name:       name,
			Username:   username,
			Password:   string(bcryptPassword),
			Department: department,
		}
		students = append(students, stu)
	}

	err = service.BatchImportStudents(c, students)
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

func ImportTeacherByExcel(c *gin.Context) {
	// 获取上传的文件
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	// 验证文件后缀（仅允许 .xlsx 格式，忽略大小写）
	filename := fileHeader.Filename
	// 将文件名转为小写，方便统一判断（如 .XLSX 会被转为 .xlsx）
	lowerFilename := strings.ToLower(filename)
	// 检查是否以 .xlsx 结尾
	if !strings.HasSuffix(lowerFilename, ".xlsx") {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "仅支持 .xlsx 格式的Excel文件",
		})
		return
	}

	// 打开文件流
	srcFile, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	defer srcFile.Close()

	// 解析Excel
	file, err := excelize.OpenReader(srcFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "解析Excel失败（文件可能损坏或格式错误）",
		})
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("关闭Excel文件失败: %v\n", err)
		}
	}()

	// 读取并打印数据
	sheetName := file.GetSheetList()[0]
	rows, err := file.GetRows(sheetName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	var teachers []entity.Teacher

	// 逐行遍历import
	for i, row := range rows {
		// 跳过标题行
		if i == 0 {
			continue
		}
		if len(row) < 4 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
		name := row[0]
		username := row[1]
		password := row[2]
		department := row[3]
		bcryptPassword, _ := util.GetPwd(password)
		tea := entity.Teacher{
			Name:       name,
			Username:   username,
			Password:   string(bcryptPassword),
			Department: department,
		}
		teachers = append(teachers, tea)
	}

	err = service.BatchImportTeachers(c, teachers)
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
