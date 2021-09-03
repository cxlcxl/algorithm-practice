package main

import (
	"math"
	"strconv"
)

func main() {

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
