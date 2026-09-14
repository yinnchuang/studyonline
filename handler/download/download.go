package download

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"studyonline/dao/redis"
	"studyonline/log"
	"studyonline/service"
	"studyonline/util"
	"time"

	"github.com/gin-gonic/gin"
	redis_ "github.com/redis/go-redis/v9"
)

// 票据类型，下载时用于校验票据与接口是否匹配，避免票据被跨类型复用
const (
	kindResource = "resource"
	kindDataset  = "dataset"
	kindHomework = "homework"
)

// 票据在 Redis 中的 key 前缀
const ticketKeyPrefix = "download_ticket:"

// 票据有效期。票据会出现在 URL 中（浏览器历史、access log、Referer），
// 窗口不能太长；但也不能一次性消费，否则浏览器对同一个 URL 的断点续传
// 和自动重试会直接失败，所以这里靠短 TTL 兜底。
const ticketTTL = 10 * time.Minute

// ticketInfo 票据内容。只存在服务端 Redis，客户端仅持有随机串，
// 并绑定票据类型、目标 id 和签发人身份，防止票据被跨资源复用。
type ticketInfo struct {
	Kind     string
	ID       uint
	UserId   uint
	Identity int
}

func (t ticketInfo) encode() string {
	return fmt.Sprintf("%s|%d|%d|%d", t.Kind, t.ID, t.UserId, t.Identity)
}

func decodeTicket(raw string) (ticketInfo, error) {
	parts := strings.Split(raw, "|")
	if len(parts) != 4 {
		return ticketInfo{}, errors.New("invalid ticket payload")
	}
	id, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return ticketInfo{}, err
	}
	userId, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return ticketInfo{}, err
	}
	identity, err := strconv.Atoi(parts[3])
	if err != nil {
		return ticketInfo{}, err
	}
	return ticketInfo{Kind: parts[0], ID: uint(id), UserId: uint(userId), Identity: identity}, nil
}

// issueTicket 生成下载票据并写入 Redis
func issueTicket(ctx context.Context, kind string, id uint, userId uint, identity int) (string, error) {
	ticket := util.GenerateToken()
	info := ticketInfo{Kind: kind, ID: id, UserId: userId, Identity: identity}
	if err := redis.RDB.Set(ctx, ticketKeyPrefix+ticket, info.encode(), ticketTTL).Err(); err != nil {
		return "", err
	}
	return ticket, nil
}

// loadTicket 读取并解析票据，票据不存在或已过期返回错误
func loadTicket(ctx context.Context, ticket string) (ticketInfo, error) {
	raw, err := redis.RDB.Get(ctx, ticketKeyPrefix+ticket).Result()
	if err != nil {
		if errors.Is(err, redis_.Nil) {
			return ticketInfo{}, errors.New("ticket expired")
		}
		return ticketInfo{}, err
	}
	return decodeTicket(raw)
}

// loadTicketFor 读取票据并校验类型，校验失败时直接写响应
func loadTicketFor(c *gin.Context, kind string) (ticketInfo, bool) {
	ticket := c.Query("ticket")
	if ticket == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "下载链接已失效，请重新获取"})
		return ticketInfo{}, false
	}
	info, err := loadTicket(c, ticket)
	if err != nil || info.Kind != kind {
		c.JSON(http.StatusBadRequest, gin.H{"message": "下载链接已失效，请重新获取"})
		return ticketInfo{}, false
	}
	return info, true
}

// buildFilename 用业务名称 + 上传文件的后缀拼出用户看到的下载文件名。
// 磁盘上的文件名是时间戳，不能直接暴露给用户。
func buildFilename(name string, filePath string) string {
	ext := filepath.Ext(filePath)
	name = strings.TrimSpace(strings.NewReplacer("/", "_", "\\", "_").Replace(name))
	if name == "" {
		return filepath.Base(filePath)
	}
	if ext == "" || strings.EqualFold(filepath.Ext(name), ext) {
		return name
	}
	return name + ext
}

// userDisplay 取下载人的姓名和部门用于下载日志；管理员没有对应记录时返回空串
func userDisplay(userId uint, identity int) (string, string) {
	info, err := service.GetUserInfo(userId, identity)
	if err != nil || info == nil {
		return "", ""
	}
	return info.Name, info.Department
}

type ResourceTicketDTO struct {
	ResourceId uint `json:"resource_id"`
}

// CreateResourceTicket 校验 token 与资源后签发一次性下载票据
func CreateResourceTicket(c *gin.Context) {
	dto := ResourceTicketDTO{}
	if err := c.ShouldBindJSON(&dto); err != nil || dto.ResourceId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求失败"})
		return
	}

	resource, err := service.GetResourceByID(c, dto.ResourceId)
	if err != nil || resource.FilePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求失败"})
		return
	}

	ticket, err := issueTicket(c, kindResource, dto.ResourceId, c.GetUint("userId"), c.GetInt("identity"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "请求失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data": gin.H{
			"url":      fmt.Sprintf("/download/resource?ticket=%s", ticket),
			"filename": buildFilename(resource.Name, resource.FilePath),
		},
	})
}

type DatasetTicketDTO struct {
	DatasetId uint `json:"dataset_id"`
}

// CreateDatasetTicket 校验 token 与数据集下载权限后签发下载票据
func CreateDatasetTicket(c *gin.Context) {
	dto := DatasetTicketDTO{}
	if err := c.ShouldBindJSON(&dto); err != nil || dto.DatasetId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求失败"})
		return
	}

	dataset, err := service.GetDatasetByID(c, dto.DatasetId)
	if err != nil || dataset.FilePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求失败"})
		return
	}

	userId := c.GetUint("userId")
	identity := c.GetInt("identity")
	if dataset.Private && !service.IfUserHasDatasetPermission(userId, identity, dto.DatasetId) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "无下载权限，请申请权限"})
		return
	}

	ticket, err := issueTicket(c, kindDataset, dto.DatasetId, userId, identity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "请求失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data": gin.H{
			"url":      fmt.Sprintf("/download/dataset?ticket=%s", ticket),
			"filename": buildFilename(dataset.Name, dataset.FilePath),
		},
	})
}

type HomeworkTicketDTO struct {
	HomeworkId uint `json:"homework_id"`
}

// CreateHomeworkTicket 校验 token 与作业后签发下载票据
func CreateHomeworkTicket(c *gin.Context) {
	dto := HomeworkTicketDTO{}
	if err := c.ShouldBindJSON(&dto); err != nil || dto.HomeworkId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求失败"})
		return
	}

	homework, err := service.GetHomeworkById(c, dto.HomeworkId)
	if err != nil || homework.FilePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求失败"})
		return
	}

	ticket, err := issueTicket(c, kindHomework, dto.HomeworkId, c.GetUint("userId"), c.GetInt("identity"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "请求失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data": gin.H{
			"url":      fmt.Sprintf("/download/homework?ticket=%s", ticket),
			"filename": buildFilename(homework.Title, homework.FilePath),
		},
	})
}

// DownloadResource 凭票据下载资源，无需携带 token
func DownloadResource(c *gin.Context) {
	info, ok := loadTicketFor(c, kindResource)
	if !ok {
		return
	}

	resource, err := service.GetResourceByID(c, info.ID)
	if err != nil || resource.FilePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求失败"})
		return
	}

	service.PlusResourceDownloadTime(c, info.ID)
	c.FileAttachment(resource.FilePath, buildFilename(resource.Name, resource.FilePath))
}

// DownloadDataset 凭票据下载数据集，无需携带 token
func DownloadDataset(c *gin.Context) {
	info, ok := loadTicketFor(c, kindDataset)
	if !ok {
		return
	}

	dataset, err := service.GetDatasetByID(c, info.ID)
	if err != nil || dataset.FilePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求失败"})
		return
	}

	// 签发时已校验过一次，这里用票据内的身份再校验一次，
	// 避免票据有效期内权限被回收后仍能下载
	if dataset.Private && !service.IfUserHasDatasetPermission(info.UserId, info.Identity, info.ID) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "无下载权限，请申请权限"})
		return
	}

	service.PlusDatasetDownloadTime(c, info.ID)

	name, department := userDisplay(info.UserId, info.Identity)
	service.AddLog(c, name, department, dataset.Name)
	log.DownloadLogger.Log(fmt.Sprintf("%s %s %s", name, department, dataset.Name))

	c.FileAttachment(dataset.FilePath, buildFilename(dataset.Name, dataset.FilePath))
}

// DownloadHomework 凭票据下载作业附件，无需携带 token
func DownloadHomework(c *gin.Context) {
	info, ok := loadTicketFor(c, kindHomework)
	if !ok {
		return
	}

	homework, err := service.GetHomeworkById(c, info.ID)
	if err != nil || homework.FilePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求失败"})
		return
	}

	c.FileAttachment(homework.FilePath, buildFilename(homework.Title, homework.FilePath))
}
