package util

import (
	"os"
	"path/filepath"
	"strings"
)

// 静态文件根目录，所有上传文件都必须落在该目录下，删除时不允许越界
const staticRoot = "static"

// IsDefaultCover 判断是否为按分类共享的系统默认封面。
// 默认封面路径形如 ./static/cover/cover1.png，被同一分类下的多个资源/数据集共用，
// 删除会导致其他记录的封面失效，因此不允许删除。
func IsDefaultCover(path string) bool {
	if path == "" {
		return false
	}
	clean := filepath.Clean(path)
	if filepath.Base(filepath.Dir(clean)) != "cover" {
		return false
	}
	base := filepath.Base(clean)
	return strings.HasPrefix(base, "cover") && strings.EqualFold(filepath.Ext(base), ".png")
}

// RemoveStaticFile 安全删除 static 目录下的单个上传文件。
// 1. 空路径、默认封面直接跳过；
// 2. 路径必须位于 static 目录内，避免路径穿越误删业务目录/系统文件；
// 3. 文件不存在视为删除成功（幂等）。
func RemoveStaticFile(path string) error {
	if path == "" || IsDefaultCover(path) {
		return nil
	}

	absPath, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		return err
	}
	absRoot, err := filepath.Abs(staticRoot)
	if err != nil {
		return err
	}

	rel, err := filepath.Rel(absRoot, absPath)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(os.PathSeparator)) {
		return &os.PathError{Op: "remove", Path: path, Err: os.ErrPermission}
	}

	if err := os.Remove(absPath); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

// RemoveStaticFiles 批量删除 static 目录下的上传文件，尽力删除全部，返回第一个错误。
func RemoveStaticFiles(paths ...string) error {
	var firstErr error
	for _, path := range paths {
		if err := RemoveStaticFile(path); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}
