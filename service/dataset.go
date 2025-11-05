package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListDatasetByTeacherId(c *gin.Context, teacherId uint) ([]entity.Dataset, error) {
	var datasets []entity.Dataset
	err := mysql.DB.Model(&entity.Dataset{}).Where("teacher_id = ?", teacherId).Find(&datasets).Error
	if err != nil {
		return nil, err
	}
	return datasets, nil
}

func ListDatasetWithLimitOffset(ctx context.Context, limit int, offset int) ([]entity.Dataset, error) {
	var datasets []entity.Dataset
	if limit == -1 {
		err := mysql.DB.Model(&entity.Dataset{}).Find(&datasets).Error
		if err != nil {
			return nil, err
		}
		return datasets, nil
	}
	err := mysql.DB.Model(&entity.Dataset{}).Order("id DESC").Limit(limit).Offset(offset).Find(&datasets).Error
	if err != nil {
		return nil, err
	}
	return datasets, nil
}

func CountDataset(ctx context.Context) (int64, error) {
	var count int64
	err := mysql.DB.Model(&entity.Dataset{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func ListDatasetWithCategoryLimitOffset(ctx context.Context, limit int, offset int, category int) ([]entity.Dataset, error) {
	var datasets []entity.Dataset
	err := mysql.DB.Model(&entity.Dataset{}).Order("id DESC").Where("category_id = ?", category).Limit(limit).Offset(offset).Find(&datasets).Error
	if err != nil {
		return nil, err
	}
	return datasets, nil
}

func CountDatasetWithCategory(ctx context.Context, category int) (int64, error) {
	var count int64
	err := mysql.DB.Model(&entity.Dataset{}).Where("category_id = ?", category).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CreateDataset(ctx context.Context, name string, categoryId int, description string, filePath string, coverPath string, scale string, teacherId uint, private bool, url string) (*entity.Dataset, error) {
	dataset := entity.Dataset{
		Name:        name,
		CategoryID:  categoryId,
		Description: description,
		Scale:       scale,
		FilePath:    filePath,
		CoverPath:   coverPath,
		TeacherId:   teacherId,
		Private:     private,
		Url:         url,
	}
	err := mysql.DB.Create(&dataset).Error
	if err != nil {
		return nil, err
	}
	return &dataset, nil
}

func UpdateDataset(ctx context.Context, datasetId uint, name string, categoryId int, description string, filePath string, coverPath string, scale string, teacherId uint, private bool, url string) error {
	dataset := entity.Dataset{
		Name:        name,
		CategoryID:  categoryId,
		Description: description,
		Scale:       scale,
		FilePath:    filePath,
		CoverPath:   coverPath,
		TeacherId:   teacherId,
		Private:     private,
		Url:         url,
	}
	err := mysql.DB.Where("id = ?", datasetId).Updates(&dataset).Error
	if err != nil {
		return err
	}
	return nil
}

func GetDatasetByID(ctx context.Context, id uint) (*entity.Dataset, error) {
	var dataset entity.Dataset
	err := mysql.DB.First(&dataset, id).Error
	if err != nil {
		return nil, err
	}
	return &dataset, nil
}

func PlusDatasetDownloadTime(ctx context.Context, datasetId uint) {
	mysql.DB.Model(&entity.Dataset{}).Where("id = ?", datasetId).
		Update("download_time", gorm.Expr("download_time + ?", 1))
}

func DeleteDataset(ctx context.Context, id uint) error {
	return mysql.DB.Delete(&entity.Dataset{}, id).Error
}

// SearchDatasetByKeyword 根据关键词搜索数据集（匹配名称或描述）
// 支持分页（limit=-1 时返回全部结果）
func SearchDatasetByKeyword(ctx context.Context, limit int, offset int, keyword string) ([]entity.Dataset, error) {
	var datasets []entity.Dataset
	// 构建基础查询：匹配名称或描述包含关键词（模糊搜索），按ID倒序（与其他列表接口一致）
	query := mysql.DB.Model(&entity.Dataset{}).
		Where("name LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%").
		Order("id DESC")

	// 处理分页逻辑（复用 limit=-1 表示查询全部的规则）
	if limit == -1 {
		err := query.Find(&datasets).Error
		if err != nil {
			return nil, err
		}
		return datasets, nil
	}

	// 带分页的查询（限制条数和偏移量）
	err := query.Limit(limit).Offset(offset).Find(&datasets).Error
	if err != nil {
		return nil, err
	}
	return datasets, nil
}

// CountDatasetByKeyword 统计符合关键词搜索条件的数据集总数
func CountDatasetByKeyword(ctx context.Context, keyword string) (int64, error) {
	var count int64
	// 与搜索方法使用完全一致的查询条件，确保总数准确
	err := mysql.DB.Model(&entity.Dataset{}).
		Where("name LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%").
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
