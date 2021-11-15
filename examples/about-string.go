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

// LongestCommonPrefix 编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。
//示例 1：
//输入：strs = ["flower","flow","flight"]
//输出："fl"
func LongestCommonPrefix(strs []string) string {
	prefix := ""
	minLen := 0
	// 取出最短子串
	for i := 0; i < len(strs); i++ {
		if strLen := len(strs[i]); i == 0 || minLen > strLen {
			minLen = strLen
		}
	}
	if minLen == 0 {
		return prefix
	}
	// 进行字符串的横向遍历
	for i := 0; i < minLen; i++ {
		for k, str := range strs {
			// 每次遍历到第一个子串重置公共前缀
			if k == 0 {
				prefix = str[0 : i+1]
				continue
			}
			// 不等时
			if str[0:i+1] != prefix {
				if k == 0 {
					return prefix
				} else {
					return prefix[0 : len(prefix)-1]
				}
			}
		}
	}

	return prefix
}

// FindLongestWord 给你一个字符串 s 和一个字符串数组 dictionary 作为字典，找出并返回字典中最长的字符串，该字符串可以通过删除 s 中的某些字符得到。
//如果答案不止一个，返回长度最长且字典序最小的字符串。如果答案不存在，则返回空字符串。
//输入：s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]
//输出："apple"
//输入：s = "abpcplea", dictionary = ["a","b","c"]
//输出："a"
func FindLongestWord(s string, dictionary []string) string {
	sByte := []byte(s)
	ret := ""
	for _, str := range dictionary {
		tmp := make([]byte, 0)
		tmp = append(tmp, sByte...)
		strByte := []byte(str)
		for i, v := range strByte {
			if byteIdx := ByteInArray(tmp, v); byteIdx != -1 {
				tmp = append(tmp[0:byteIdx], tmp[byteIdx+1:]...)
			} else {
				break
			}
			if i == len(strByte)-1 && len(ret) < len(strByte) {
				ret = str
			}
		}
	}
	return ret
}

//给你一个由若干括号和字母组成的字符串 s ，删除最小数量的无效括号，使得输入的字符串有效。
//返回所有可能的结果。答案可以按 任意顺序 返回。
//输入：s = "()())()"
//输出：["(())()","()()()"]
//输入：s = "(a)())()"
//输出：["(a())()","(a)()()"]
//输入：s = ")("
//输出：[""]
func removeInvalidParentheses(s string) []string {
	return []string{}
}
