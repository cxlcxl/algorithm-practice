package examples

import (
	"strings"
)

// ContinueStrings 给定一个字符串 s ，请你找出其中不含有重复字符的 连续最长子串 的长度。
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
func ContinueStrings(str string) int {
	maxLen := 0
	// 改成字节后内存消耗更少
	hasStrings := make([]byte, 0)
	for i := 0; i < len(str); i++ {
		current := []byte(str[i : i+1])
		if key := ByteInArray(hasStrings, current[0]); key != -1 {
			if stringsLen := len(hasStrings); maxLen < stringsLen {
				maxLen = stringsLen
			}
			hasStrings = hasStrings[key+1:]
		}
		hasStrings = append(hasStrings, current[0])
	}
	if stringsLen := len(hasStrings); stringsLen > maxLen {
		return stringsLen
	}
	return maxLen
}

// IsValid 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
func IsValid(s string) bool {
	if n := len(s); n%2 == 0 {
		for strings.Contains(s, "()") || strings.Contains(s, "[]") || strings.Contains(s, "{}") {
			s = strings.ReplaceAll(s, "()", "")
			s = strings.ReplaceAll(s, "[]", "")
			s = strings.ReplaceAll(s, "{}", "")
		}
		if len(s) == 0 {
			return true
		}
	}
	return false
}
