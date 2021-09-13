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
