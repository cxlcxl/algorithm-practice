package study_plan

import (
	"math"
	"sort"
)

// ArrayHasRepeat 给定一个整数数组，判断是否存在重复元素。
//如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
func ArrayHasRepeat(nums []int) bool {
	//tmp := map[int]int{}
	//for _, num := range nums {
	//	if _, ok := tmp[num]; !ok {
	//		tmp[num] = 0
	//	} else {
	//		return true
	//	}
	//}
	//return false

	for i := 0; i < len(nums); i++ {
		for k := i + 1; k < len(nums); k++ {
			if nums[i] == nums[k] {
				return true
			}
		}
	}
	return false
}

// MaxSubArray 给定一个整数数组 nums，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
func MaxSubArray(nums []int) int {
	x, tmp := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		// 前面的和小于等于当前数字，舍弃前面的，以当前值开始加
		if tmp+nums[i] <= nums[i] {
			tmp = nums[i]
		} else {
			tmp += nums[i]
		}
		// 更新临时的最大值
		if x < tmp {
			x = tmp
		}
	}
	return x
}

// AddTwoNums 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
func AddTwoNums(nums []int, target int) []int {
	mapNums := map[int]int{}
	for i := 0; i < len(nums); i++ {
		key := target - nums[i]
		if idx, ok := mapNums[key]; ok {
			return []int{i, idx}
		} else {
			mapNums[nums[i]] = i
		}
	}
	return []int{-1, -1}
}

// MergeArrays 给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
//注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
//输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//输出：[1,2,2,3,5,6]
//解释：需要合并 [1,2,3] 和 [2,5,6]
//合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素
//注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中
func MergeArrays(numsx []int, x int, numsy []int, y int) []int {
	numsx = append(numsx[0:x], numsy...)
	sort.Ints(numsx)
	return numsx
}

// ArrayIntersect 给定两个数组，编写一个函数来计算它们的交集。
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[4,9]
func ArrayIntersect(nums1 []int, nums2 []int) []int {
	intersect := make([]int, 0)
	sort.Ints(nums1)
	for i := 0; i < len(nums2); i++ {
		if set := intersectInArray(nums1, nums2[i]); set != -1 {
			nums1 = append(nums1[0:set], nums1[set+1:]...)
			intersect = append(intersect, nums2[i])
		}
	}
	return intersect
}
func intersectInArray(x []int, target int) int {
	for key, v := range x {
		if target < v {
			break
		}
		if v == target {
			return key
		}
	}
	return -1
}

// hashmap方式
func arrayIntersectMap(nums1 []int, nums2 []int) []int {
	tmp := map[int]int{}
	intersect := make([]int, 0)
	for _, num := range nums1 {
		if _, ok := tmp[num]; ok {
			tmp[num]++
		} else {
			tmp[num] = 1
		}
	}
	for _, num := range nums2 {
		if v, ok := tmp[num]; ok && v > 0 {
			intersect = append(intersect, num)
			tmp[num]--
		}
	}
	return intersect
}

// MaxProfit 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
//你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
//输入：[2,8,1,5,3,6,4]
//输出：5
//解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
//输入：prices = [7,6,4,3,1]
//输出：0
//解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
func MaxProfit(prices []int) int {
	min, max := math.MaxInt32, 0
	for _, v := range prices {
		if v < min {
			// 找出历史最低
			min = v
		} else if v-min > max {
			// 后面的
			max = v - min
		}
	}

	return max
}

// MatrixReshape 在 MATLAB 中，有一个非常有用的函数 reshape ，它可以将一个 m x n 矩阵重塑为另一个大小不同（r x c）的新矩阵，但保留其原始数据。
//给你一个由二维数组 mat 表示的 m x n 矩阵，以及两个正整数 r 和 c ，分别表示想要的重构的矩阵的行数和列数。
//重构后的矩阵需要将原始矩阵的所有元素以相同的 行遍历顺序 填充。
//如果具有给定参数的 reshape 操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。
//输入：mat = [[1,2],[3,4]], r = 1, c = 4
//输出：[[1,2,3,4]]
func MatrixReshape(mat [][]int, r int, c int) [][]int {
	matTotal := 0
	tmp := make([]int, 0)
	for _, nums := range mat {
		matTotal += len(nums)
		tmp = append(tmp, nums...)
	}
	if r*c != matTotal {
		return mat
	}
	ret := make([][]int, 0)
	list := 0
	listInt := make([]int, c)
	for _, num := range tmp {
		listInt[list] = num
		list++
		if list >= c {
			ret = append(ret, listInt)
			listInt = make([]int, c)
			list = 0
		}
	}
	return ret
}

// Generate 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//在「杨辉三角」中，每个数是它左上方和右上方的数的和。
//输入: numRows = 5
//输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
//输入: numRows = 1
//输出: [[1]]
func Generate(numRows int) [][]int {
	rows := make([][]int, numRows)
	if numRows == 0 {
		return rows
	}
	rows[0] = []int{1}
	if numRows == 1 {
		return rows
	}
	rows[1] = []int{1, 1}
	for i := 2; i < numRows; i++ {
		item := make([]int, i+1)
		for k := 0; k < i+1; k++ {
			if k == 0 || k == i {
				item[k] = 1
			} else {
				item[k] = rows[i-1][k] + rows[i-1][k-1]
			}
		}
		rows[i] = item
	}

	return rows
}

// FindPeakElement 寻找峰值
//峰值元素是指其值严格大于左右相邻值的元素。
//给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
//你可以假设 nums[-1] = nums[n] = -∞ 。
//你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
func FindPeakElement(nums []int) int {
	peak := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			peak = i
		} else if nums[i] < nums[i-1] {
			break
		}
	}
	return peak
}

// CanConstruct 给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，判断第一个字符串 ransom 能不能由第二个字符串 magazines 里面的字符构成。如果可以构成，返回 true ；否则返回 false。
//(题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。杂志字符串中的每个字符只能在赎金信字符串中使用一次。)
func CanConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	magazineBytes, ransomNoteBytes := []byte(magazine), []byte(ransomNote)
	magazineMap := make(map[byte]int, 0)
	for _, magazineByte := range magazineBytes {
		if _, ok := magazineMap[magazineByte]; ok {
			magazineMap[magazineByte]++
		} else {
			magazineMap[magazineByte] = 1
		}
	}
	for _, noteByte := range ransomNoteBytes {
		if v, ok := magazineMap[noteByte]; ok && v > 0 {
			magazineMap[noteByte]--
		} else {
			return false
		}
	}
	return true
}

// FirstUniqChar 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
//s = "leetcode"
//返回 0
//s = "loveleetcode"
//返回 2
func FirstUniqChar(s string) int {
	sMap := make(map[byte]int, 0)
	bytes := []byte(s)
	for _, b := range bytes {
		sMap[b]++
	}
	for i, b := range bytes {
		if sMap[b] == 1 {
			return i
		}
	}
	return -1
}

// SetZeroes 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
//输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
//输出：[[1,0,1],[0,0,0],[1,0,1]]
func SetZeroes(matrix [][]int) [][]int {
	zeroX := make(map[int]int, 0)
	zeroY := make(map[int]int, 0)
	for i, ints := range matrix {
		for i2, i3 := range ints {
			if i3 == 0 {
				zeroY[i2] = 0
				zeroX[i] = 0
			}
		}
	}
	for i, ints := range matrix {
		for i2, _ := range ints {
			if _, ok := zeroX[i]; ok {
				ints[i2] = 0
			}
			if _, ok := zeroY[i2]; ok {
				ints[i2] = 0
			}
		}
	}

	return matrix
}

// IsAnagram 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
//输入: s = "anagram", t = "nagaram"
//输出: true
//输入: s = "rat", t = "car"
//输出: false
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[byte]int, len(s))
	tMap := make(map[byte]int, len(s))

	bytes := []byte(s)
	bytet := []byte(t)
	for i := 0; i < len(bytes); i++ {
		sMap[bytes[i]] += 1
		tMap[bytet[i]] += 1
	}
	for _, b := range bytes {
		if v, ok := tMap[b]; !ok || v != sMap[b] {
			return false
		}
	}

	return true
}

// CanWinNim 你和你的朋友，两个人一起玩 Nim 游戏：
//桌子上有一堆石头。
//你们轮流进行自己的回合，你作为先手。
//每一回合，轮到的人拿掉 1 - 3 块石头。
//拿掉最后一块石头的人就是获胜者。
//假设你们每一步都是最优解。请编写一个函数，来判断你是否可以在给定石头数量为 n 的情况下赢得游戏。如果可以赢，返回 true；否则，返回 false 。
//输入：n = 4
//输出：false
//解释：如果堆中有 4 块石头，那么你永远不会赢得比赛；
//因为无论你拿走 1 块、2 块 还是 3 块石头，最后一块石头总是会被你的朋友拿走。
//输入：n = 1
//输出：true
//输入：n = 2
//输出：true
func CanWinNim(n int) bool {
	return n%4 != 0
}

// IntToRoman 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。
//通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。
//同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
//给你一个整数，将其转为罗马数字。
//输入: num = 3
//输出: "III"
//输入: num = 4
//输出: "IV"
//输入: num = 9
//输出: "IX"
//输入: num = 58
//输出: "LVIII"
//解释: L = 50, V = 5, III = 3.
//输入: num = 1994
//输出: "MCMXCIV"
//解释: M = 1000, CM = 900, XC = 90, IV = 4.
func IntToRoman(num int) string {
	romanMap := map[int]byte{
		1:    'I',
		5:    'V',
		10:   'X',
		50:   'L',
		100:  'C',
		500:  'D',
		1000: 'M',
	}
	roman := make([]byte, 0)
	i := 1
	for num > 0 {
		tmp := num % 10
		num /= 10
		switch tmp {
		case 1:
			roman = append(roman, romanMap[i])
		case 2:
			roman = append(roman, romanMap[i], romanMap[i])
		case 3:
			roman = append(roman, romanMap[i], romanMap[i], romanMap[i])
		case 4:
			roman = append(roman, romanMap[i*5], romanMap[i])
		case 5:
			roman = append(roman, romanMap[i*5])
		case 6:
			roman = append(roman, romanMap[i], romanMap[i*5])
		case 7:
			roman = append(roman, romanMap[i], romanMap[i], romanMap[i*5])
		case 8:
			roman = append(roman, romanMap[i], romanMap[i], romanMap[i], romanMap[i*5])
		case 9:
			roman = append(roman, romanMap[i*10], romanMap[i])
		default:
		}

		i *= 10
	}
	retBytes := make([]byte, len(roman))
	j := 0
	for k := len(roman) - 1; k >= 0; k-- {
		retBytes[j] = roman[k]
		j++
	}
	return string(retBytes)
}
