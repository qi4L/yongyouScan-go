package utils

import "strings"

// InsertBackslashX 在每两个字符后以及首字符前插入"\x"
func InsertBackslashX(s string) string {
	var result strings.Builder
	result.WriteString("\\x") // 在开始时先写入"\x"

	for i, r := range s {
		if i > 0 && i%2 == 0 {
			result.WriteString("\\x") // 每两个字符后插入"\x"
		}
		result.WriteRune(r) // 写入当前字符
	}
	return result.String()
}
