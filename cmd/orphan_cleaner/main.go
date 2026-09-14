// Command orphan_cleaner 清理 static 目录下已无数据库记录引用的孤儿文件。
//
// 背景：资源/数据集/作业/提交记录均为软删除，早期版本只删数据库条目、不删本地文件，
// 导致磁盘被历史文件占满。本脚本以数据库中未删除记录引用的路径作为白名单，
// 扫描上传目录，列出（或删除）不在白名单中的文件。
//
// 用法（必须在项目根目录运行，脚本依赖 ./init/project.ini 与相对路径 ./static）：
//
//	go run ./cmd/orphan_cleaner                  # 干跑，只输出清单，不删除
//	go run ./cmd/orphan_cleaner -apply           # 确认清单无误后真正删除
//	go run ./cmd/orphan_cleaner -min-age=72h     # 只处理 72 小时前上传的文件
//
// 说明：默认封面 static/cover/cover{N}.png 按分类共享，永不清理。
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"studyonline/dao/entity"
	"studyonline/dao/mysql"
	"studyonline/util"

	"gorm.io/gorm"
)

// 上传文件存放的子目录，只在这些目录内扫描，避免误删 static 下的静态资源（如 format.xlsx）
var uploadedDirs = []string{"resource", "dataset", "cover", "homework"}

const previewCount = 20

func main() {
	apply := flag.Bool("apply", false, "真正删除孤儿文件；默认只输出清单")
	minAge := flag.Duration("min-age", 24*time.Hour,
		"只清理最后修改时间早于该时长的文件，避免误删刚上传、尚未入库的文件")
	report := flag.String("report", "orphan_files.txt", "孤儿文件清单输出路径")
	flag.Parse()

	db, err := mysql.Open()
	if err != nil {
		fmt.Println("数据库连接失败，请检查 ./init/project.ini:", err)
		os.Exit(1)
	}

	referenced, err := loadReferencedFiles(db)
	if err != nil {
		fmt.Println("加载数据库白名单失败:", err)
		os.Exit(1)
	}
	fmt.Printf("数据库中被引用的文件数: %d\n", len(referenced))

	orphans, totalSize, err := findOrphanFiles(referenced, *minAge)
	if err != nil {
		fmt.Println("扫描上传目录失败:", err)
		os.Exit(1)
	}
	fmt.Printf("发现孤儿文件: %d 个，共 %.2f MB\n", len(orphans), float64(totalSize)/1024/1024)
	for i, path := range orphans {
		if i >= previewCount {
			fmt.Printf("... 其余 %d 个见 %s\n", len(orphans)-previewCount, *report)
			break
		}
		fmt.Println(" ", path)
	}

	if err := writeReport(*report, orphans); err != nil {
		fmt.Println("写入清单失败:", err)
		os.Exit(1)
	}

	if !*apply {
		fmt.Println("当前为干跑模式，未删除任何文件。确认清单无误后加 -apply 执行删除（建议先备份 static 目录）")
		return
	}

	var deleted int
	var freed int64
	for _, path := range orphans {
		info, err := os.Stat(path)
		if err == nil {
			freed += info.Size()
		}
		if err := util.RemoveStaticFile(path); err != nil {
			fmt.Printf("删除失败: %s: %v\n", path, err)
			continue
		}
		deleted++
	}
	fmt.Printf("删除完成: %d/%d 个文件，释放 %.2f MB\n", deleted, len(orphans), float64(freed)/1024/1024)
}

// loadReferencedFiles 汇总各业务表中未删除记录引用的文件路径作为白名单。
// gorm 查询默认带 deleted_at IS NULL，因此已软删除记录对应的文件即孤儿文件。
func loadReferencedFiles(db *gorm.DB) (map[string]struct{}, error) {
	referenced := make(map[string]struct{})
	add := func(paths []string) {
		for _, path := range paths {
			path = strings.TrimSpace(path)
			if path == "" {
				continue
			}
			referenced[normalize(path)] = struct{}{}
		}
	}

	for _, item := range []struct {
		model  interface{}
		column string
	}{
		{&entity.Resource{}, "file_path"},
		{&entity.Resource{}, "cover_path"},
		{&entity.Dataset{}, "file_path"},
		{&entity.Dataset{}, "cover_path"},
		{&entity.Homework{}, "file_path"},
		{&entity.Submission{}, "file_path"},
	} {
		var paths []string
		if err := db.Model(item.model).Pluck(item.column, &paths).Error; err != nil {
			return nil, fmt.Errorf("读取 %T.%s 失败: %w", item.model, item.column, err)
		}
		add(paths)
	}
	return referenced, nil
}

func findOrphanFiles(referenced map[string]struct{}, minAge time.Duration) ([]string, int64, error) {
	var orphans []string
	var totalSize int64
	cutoff := time.Now().Add(-minAge)

	for _, dir := range uploadedDirs {
		root := filepath.Join("static", dir)
		err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				if os.IsNotExist(err) {
					return nil
				}
				return err
			}
			if entry.IsDir() {
				return nil
			}
			// 按分类共享的默认封面，多个记录共用，不参与清理
			if util.IsDefaultCover(path) {
				return nil
			}
			if _, ok := referenced[normalize(path)]; ok {
				return nil
			}
			info, err := entry.Info()
			if err != nil {
				return nil
			}
			if info.ModTime().After(cutoff) {
				fmt.Printf("跳过(近期上传，可能尚未入库): %s\n", path)
				return nil
			}
			orphans = append(orphans, path)
			totalSize += info.Size()
			return nil
		})
		if err != nil {
			return nil, 0, err
		}
	}
	return orphans, totalSize, nil
}

func writeReport(path string, orphans []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, orphan := range orphans {
		if _, err := file.WriteString(orphan + "\n"); err != nil {
			return err
		}
	}
	return nil
}

// normalize 统一数据库存储路径("./static/...")与扫描路径("static/...")的表现形式
func normalize(path string) string {
	return strings.TrimPrefix(filepath.Clean(path), "."+string(os.PathSeparator))
}
