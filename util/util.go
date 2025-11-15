package util

import (
	"unicode/utf8"
)

// ProtectName 将姓名除了第一个字符外的所有字符替换为*
func ProtectName(name string) string {
	if name == "" {
		return ""
	}

	// 获取第一个字符
	firstChar, size := utf8.DecodeRuneInString(name)
	if firstChar == utf8.RuneError {
		return name
	}

	// 如果只有一个字符，直接返回
	if len(name) == size {
		return name
	}

	// 构建结果：第一个字符 + 对应数量的*
	result := string(firstChar)
	for i := size; i < len(name); {
		_, nextSize := utf8.DecodeRuneInString(name[i:])
		result += "*"
		i += nextSize
	}

	return result
}
