package examples

import (
	"fmt"
	"math"
	"math/bits"
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

// IsPowerOfThree 给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。
//整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x
//输入：n = 27
//输出：true
//输入：n = 0
//输出：false
//输入：n = 9
//输出：true
//输入：n = 45
//输出：false
func IsPowerOfThree(n int) bool {
	for n >= 3 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return n == 1
}

// CountAndSay 给定一个正整数 n ，输出外观数列的第 n 项。
//「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。
//你可以将其视作是由递归公式定义的数字字符串序列：
//countAndSay(1) = "1"
//countAndSay(n) 是对 countAndSay(n-1) 的描述，然后转换成另一个数字字符串。
//前五项如下：
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//第一项是数字 1
//描述前一项，这个数是 1 即 “ 一 个 1 ”，记作 "11"
//描述前一项，这个数是 11 即 “ 二 个 1 ” ，记作 "21"
//描述前一项，这个数是 21 即 “ 一 个 2 + 一 个 1 ” ，记作 "1211"
//描述前一项，这个数是 1211 即 “ 一 个 1 + 一 个 2 + 二 个 1 ” ，记作 "111221"
//要 描述 一个数字字符串，首先要将字符串分割为 最小 数量的组，每个组都由连续的最多 相同字符 组成。
//然后对于每个组，先描述字符的数量，然后描述字符，形成一个描述组。
//要将描述转换为数字字符串，先将每组中的字符数量用数字替换，再将所有描述组连接起来。
//输入：n = 1
//输出："1"
//解释：这是一个基本样例。
//输入：n = 4
//输出："1211"
//解释：
//countAndSay(1) = "1"
//countAndSay(2) = 读 "1" = 一 个 1 = "11"
//countAndSay(3) = 读 "11" = 二 个 1 = "21"
//countAndSay(4) = 读 "21" = 一 个 2 + 一 个 1 = "12" + "11" = "1211"
func CountAndSay(n int) string {
	dict := "1"
	for i := 2; i <= n; i++ {
		str, tmpString := "", ""
		tmpNum := 0
		for k := 0; k < len(dict); k++ {
			if tmpString == "" {
				tmpString = dict[k : k+1]
				tmpNum++
				continue
			}
			if tmpString == dict[k:k+1] {
				tmpNum++
			} else {
				str += strconv.Itoa(tmpNum) + tmpString
				tmpString = dict[k : k+1]
				tmpNum = 1
			}
		}
		str += strconv.Itoa(tmpNum) + tmpString
		dict = str
	}
	return dict
}

// FindComplement 给你一个 正 整数 num ，输出它的补数。补数是对该数的二进制表示取反。
//输入：num = 5
//输出：2
//解释：5 的二进制表示为 101（没有前导零位），其补数为 010。所以你需要输出 2 。
//输入：num = 1
//输出：0
//解释：1 的二进制表示为 1（没有前导零位），其补数为 0。所以你需要输出 0 。
func FindComplement(num int) int {
	//n := strconv.FormatInt(int64(num), 2)
	//dict := ""
	//for i := 0; i < len(n); i++ {
	//	if n[i:i+1] == "1" {
	//		dict += "0"
	//	} else {
	//		dict += "1"
	//	}
	//}
	//i, err := strconv.ParseInt(dict, 2, 64)
	//if err != nil {
	//	return 0
	//}
	fmt.Println(1 << bits.Len(uint(num)))
	//return int(i)
	return 1
}
