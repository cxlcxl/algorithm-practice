package examples

import (
	"math"
	"strconv"
)

// IsPalindrome 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
func IsPalindrome(x int) bool {
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

// Reverse 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//如果反转后整数超过 32 位的有符号整数的范围 [−2 的 31 次方,  （2 的 31 次方） − 1] ，就返回 0。
//假设环境不允许存储 64 位整数（有符号或无符号）。
//fmt.Println(reverse(-36469))
func Reverse(x int) (rev int) {
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
