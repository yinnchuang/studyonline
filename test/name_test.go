package test

import (
	"fmt"
	"studyonline/util"
	"testing"
)

func TestProtectName(t *testing.T) {
	// 测试用例
	testNames := []string{
		"张三",
		"李四",
		"王五",
		"诸葛亮",
		"欧阳锋",
		"慕容复",
		"A",
		"AB",
		"",
		"约翰·史密斯",
	}

	fmt.Println("姓名隐私保护工具测试结果：")
	for _, name := range testNames {
		protected := util.ProtectName(name)
		fmt.Printf("原始姓名: %-10s -> 保护后: %s\n", name, protected)
	}
}
