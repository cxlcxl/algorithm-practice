package main

import (
	"math"
	"strconv"
	"strings"
)

func main() {

}

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//fmt.Println(searchInsert([]int{1, 2, 3, 5, 9}, 10))
func searchInsert(nums []int, target int) int {
	idx := 0
	for idx <= len(nums)-1 {
		if nums[idx] >= target {
			return idx
		}
		idx++
	}
	return idx
}

//二分查找：给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target，写一个函数搜索 nums 的 target，如果目标值存在返回下标，否则返回 -1
// fmt.Println(search([]int{1, 3, 6, 7, 9, 10, 12, 15, 16, 18, 19, 22, 25, 28, 29}, 6))
func search(nums []int, target int) int {
	minIdx := 0
	maxIdx := len(nums) - 1
	if target > nums[maxIdx] || target < nums[minIdx] {
		return -1
	}
	for minIdx <= maxIdx {
		middleIdx := (maxIdx-minIdx)/2 + minIdx
		if nums[middleIdx] == target {
			return middleIdx
		}
		if nums[middleIdx] > target {
			maxIdx = middleIdx - 1
		}
		if nums[middleIdx] < target {
			minIdx = middleIdx + 1
		}
	}

	return -1
}

//输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
//例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//l2 := []int{2, 3, 3, 5, 7, 9, 12, 15, 18, 19}
//l1 := []int{1, 6, 7, 8, 10}
//l3 := mergeTwoLists(l1, l2)
//fmt.Println(l3)
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

//给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//如果反转后整数超过 32 位的有符号整数的范围 [−2 的 31 次方,  （2 的 31 次方） − 1] ，就返回 0。
//假设环境不允许存储 64 位整数（有符号或无符号）。
//fmt.Println(reverse(-36469))
func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。
// x, y := doSum([]int{1, 5, 7, 4, 6, 12, 8}, 55)
// fmt.Println(x, y)
func doSum(nums []int, sum int) (indexX, indexY int) {
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		for k := i + 1; k < numsLen; k++ {
			if nums[i]+nums[k] == sum {
				indexX = i
				indexY = k
				return
			}
		}
	}
	return -1, -1
}
