package main

import (
	"math"
	"strconv"
	"strings"
)

func main() {
	//l2 := []int{2, 3, 3, 5, 7, 9, 12, 15, 18, 19}
	//l1 := []int{1, 6, 7, 8, 10}
	//l3 := mergeTwoLists(l1, l2)
	//fmt.Println(l3)
}

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
func mergeTwoLists(l1, l2 []int) []int {
	l3 := make([]int, 0)
	for len(l1) > 0 || len(l2) > 0 {
		len1 := len(l1)
		len2 := len(l2)
		if len1 > 0 && len2 > 0 {
			if l1[0] > l2[0] {
				l3 = append(l3, l2[0])
				l2 = l2[1:]
			} else {
				l3 = append(l3, l1[0])
				l1 = l1[1:]
			}
			continue
		} else if len1 > 0 && len2 <= 0 {
			l3 = append(l3, l1[0])
			l1 = l1[1:]
			continue
		} else if len1 <= 0 && len2 > 0 {
			l3 = append(l3, l2[0])
			l2 = l2[1:]
			continue
		}
	}
	return l3
}

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
func isValid(s string) bool {
	if n := len(s); n%2 == 0 {
		// maps := map[string]int{"(": 1, ")": -1, "{": 2, "}": -2, "[": 3, "]": -3}
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

//给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if math.MaxInt32 < x || math.MinInt32 > x {
		return false
	}
	palindromeNum := ""
	y := strconv.Itoa(x)
	for i := len(y) - 1; i >= 0; i-- {
		palindromeNum += y[i : i+1]
	}
	return palindromeNum == y
}
