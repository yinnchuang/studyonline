package lessonplan

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"studyonline/dao/entity"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

var (
	// 转发代理
	target, _     = url.Parse("http://127.0.0.1:12010")
	generateProxy = httputil.NewSingleHostReverseProxy(target)
)

func GenerateLessonPlan(c *gin.Context) {
	// 读出原始 body
	body, _ := io.ReadAll(c.Request.Body)

	// 解析成 map
	var data map[string]interface{}
	_ = json.Unmarshal(body, &data)

	unitIds := data["unit_ids"].([]interface{})
	unitNames := make([]string, 0)
	for _, unitId := range unitIds {
		unit, err := service.GetUnitById(c, uint(unitId.(float64)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "请求失败",
			})
			return
		}
		unitNames = append(unitNames, unit.UnitName)
	}

	data["unit_names"] = unitNames

	// 重新编码
	newBody, _ := json.Marshal(data)

	// 5. 写回 Request
	c.Request.Body = io.NopCloser(bytes.NewReader(newBody))
	c.Request.ContentLength = int64(len(newBody))
	c.Request.Header.Set("Content-Length", strconv.Itoa(len(newBody)))

	// 6. 转发
	generateProxy.ServeHTTP(c.Writer, c.Request)
}

func GetAllLessonPlan(c *gin.Context) {
	lessonPlans, err := service.GetAllLessonPlan()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    lessonPlans,
	})
}

type RemoveLessonPlanDTO struct {
	LessonPlanId uint `json:"lesson_plan_id"`
}

func RemoveLessonPlan(c *gin.Context) {
	removeLessonPlanDTO := RemoveLessonPlanDTO{}
	err := c.ShouldBindBodyWithJSON(&removeLessonPlanDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	err = service.RemoveLessonPlan(c, removeLessonPlanDTO.LessonPlanId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	err = service.RemoveLessonPlanStudent(c, removeLessonPlanDTO.LessonPlanId)
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

type UpdateLessonPlanDTO struct {
	ID                uint   `json:"id"`
	Objectives        string `json:"objectives"`
	KeyPoints         string `json:"key_points"`
	DifficultPoints   string `json:"difficult_points"`
	Content           string `json:"content"`
	IdeologicalPoints string `json:"ideological_points"`
}

func UpdateLessonPlan(c *gin.Context) {
	updateLessonPlanDTO := UpdateLessonPlanDTO{}
	err := c.ShouldBindBodyWithJSON(&updateLessonPlanDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
			"err":     err.Error(),
		})
		return
	}
	lp, err := service.GetLessonPlanById(c, updateLessonPlanDTO.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
			"err":     err.Error(),
		})
		return
	}
	// 更新数据
	lp.Objectives = updateLessonPlanDTO.Objectives
	lp.KeyPoints = updateLessonPlanDTO.KeyPoints
	lp.DifficultPoints = updateLessonPlanDTO.DifficultPoints
	lp.Content = updateLessonPlanDTO.Content
	lp.IdeologicalPoints = updateLessonPlanDTO.IdeologicalPoints

	err = service.UpdateLessonPlan(c, lp.ID, lp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
			"err":     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}

type PublishLessonPlanDTO struct {
	ID uint `json:"id"`
}

func PublishLessonPlan(c *gin.Context) {
	publishLessonPlanDTO := PublishLessonPlanDTO{}
	err := c.ShouldBindBodyWithJSON(&publishLessonPlanDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
			"err":     err.Error(),
		})
		return
	}
	lp, err := service.GetLessonPlanById(c, publishLessonPlanDTO.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
			"err":     err.Error(),
		})
		return
	}
	// 从未发布
	if lp.PublishStatus == 0 {
		// 更新数据
		lp.PublishStatus = 1
		err = service.UpdateLessonPlan(c, lp.ID, lp)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
				"err":     err.Error(),
			})
			return
		}
	} else { // 已发布，教师侧可以不用管了，学生侧进行更新(直接删了重写)
		err := service.RemoveLessonPlanStudent(c, lp.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
				"err":     err.Error(),
			})
			return
		}

	}
	lps := entity.LessonPlanStudent{
		Title:             lp.Title,
		Duration:          lp.Duration,
		Objectives:        lp.Objectives,
		KeyPoints:         lp.KeyPoints,
		DifficultPoints:   lp.DifficultPoints,
		Content:           lp.Content,
		IdeologicalPoints: lp.IdeologicalPoints,
		UnitIds:           lp.UnitIds,
		FatherId:          lp.ID,
	}
	err = service.CreateLessonPlanStudent(c, &lps)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
			"err":     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}
