package examples

func SortInt(x []int) []int {
	tmp := 0
	for i := 0; i < len(x); i++ {
		for k := 1; k < len(x); k++ {
			if x[k] < x[k-1] {
				tmp = x[k]
				x[k] = x[k-1]
				x[k-1] = tmp
			}
		}
	}
	return x
}

// AddNums ...
//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
func AddNums(x, y []int) []int {
	xl, yl := len(x), len(y)
	xTotal, yTotal := 0, 0
	for i := xl - 1; i >= 0; i-- {
		xTotal = xTotal*10 + x[i]
	}
	for i := yl - 1; i >= 0; i-- {
		yTotal = yTotal*10 + y[i]
	}
	total := xTotal + yTotal
	rs := make([]int, 0)
	for total > 0 {
		rs = append(rs, total%10)
		total /= 10
	}
	return rs
}

// ArrayUnique ...
//给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
//不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//思路：定义一个可以截断的位置，遍历数组时发现重复的跳过，把不重复的数字赋值到截断的位置，再将截断位置 +1
//fmt.Println(ArrayUnique([]int{0, 0, 0, 1, 1, 2, 3, 3}))
func ArrayUnique(array []int) []int {
	uniqueIdx := 1
	for i := 1; i < len(array); i++ {
		if array[i] != array[i-1] {
			array[uniqueIdx] = array[i]
			uniqueIdx += 1
		}
	}
	return array[0:uniqueIdx]
}

// SearchInsert 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//fmt.Println(searchInsert([]int{1, 2, 3, 5, 9}, 10))
func SearchInsert(nums []int, target int) int {
	idx := 0
	for idx <= len(nums)-1 {
		if nums[idx] >= target {
			return idx
		}
		idx++
	}
	return idx
}

// Search 二分查找：给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target，写一个函数搜索 nums 的 target，如果目标值存在返回下标，否则返回 -1
//思路：左右两个指针，对比后移动指针
// fmt.Println(search([]int{1, 3, 6, 7, 9, 10, 12, 15, 16, 18, 19, 22, 25, 28, 29}, 6))
func Search(nums []int, target int) int {
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

// MergeTwoLists 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//l2 := []int{2, 3, 3, 5, 7, 9, 12, 15, 18, 19}
//l1 := []int{1, 6, 7, 8, 10}
//l3 := mergeTwoLists(l1, l2)
//fmt.Println(l3)
func MergeTwoLists(l1, l2 []int) []int {
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

// DoSum 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。
// x, y := doSum([]int{1, 5, 7, 4, 6, 12, 8}, 55)
// fmt.Println(x, y)
func DoSum(nums []int, sum int) (indexX, indexY int) {
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

// IntInArray ...
func IntInArray(x []int, target int) int {
	for key, v := range x {
		if v == target {
			return key
		}
	}
	return -1
}

// StringInArray ...
func StringInArray(x []string, target string) int {
	for key, v := range x {
		if v == target {
			return key
		}
	}
	return -1
}

// ByteInArray ...
func ByteInArray(x []byte, target byte) int {
	for key, v := range x {
		if v == target {
			return key
		}
	}
	return -1
}

// FindMaximizedCapital 假设 力扣（LeetCode）即将开始 IPO 。为了以更高的价格将股票卖给风险投资公司，
//力扣 希望在 IPO 之前开展一些项目以增加其资本。 由于资源有限，它只能在 IPO 之前完成最多 k 个不同的项目。
//帮助 力扣 设计完成最多 k 个不同项目后得到最大总资本的方式。
//给你 n 个项目。对于每个项目 i ，它都有一个纯利润 profits[i] ，和启动该项目需要的最小资本 capital[i] 。
//最初，你的资本为 w 。当你完成一个项目时，你将获得纯利润，且利润将被添加到你的总资本中。
//总而言之，从给定项目中选择 最多 k 个不同项目的列表，以 最大化最终资本 ，并输出最终可获得的最多资本。
//答案保证在 32 位有符号整数范围内。
//示例 1：
//输入：k = 2, w = 0, profits = [1,2,3], capital = [0,1,1]
//输出：4
//解释：
//由于你的初始资本为 0，你仅可以从 0 号项目开始。
//在完成后，你将获得 1 的利润，你的总资本将变为 1。
//此时你可以选择开始 1 号或 2 号项目。
//由于你最多可以选择两个项目，所以你需要完成 2 号项目以获得最大的资本。
//因此，输出最后最大化的资本，为 0 + 1 + 3 = 4。
//示例 2：
//输入：k = 3, w = 0, profits = [1,2,3], capital = [0,1,2]
//输出：6
func FindMaximizedCapital(k int, w int, profits []int, capital []int) int {
	// 筛选的次数限制
	profitsLen := len(profits)
	usedIdx := make([]int, profitsLen)
	for i := 0; i < profitsLen; i++ {
		usedIdx[i] = -1
	}
	i := 0
	for i < k && i < profitsLen {
		current := 0
		for j, c := range capital {
			// 未被使用 && 小于等于初始资本 && 较大的利润
			if c <= w && profits[j] > current {
				if idx := IntInArray(usedIdx, j); idx == -1 {
					usedIdx[i] = j
					current = profits[j]
				}
			}
		}
		i++
		w += current
	}
	return w
}
