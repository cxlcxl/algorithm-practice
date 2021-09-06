package examples

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
