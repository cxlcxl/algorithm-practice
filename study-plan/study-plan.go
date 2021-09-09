package study_plan

import "sort"

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
