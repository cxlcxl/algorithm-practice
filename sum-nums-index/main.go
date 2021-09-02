package main

import "fmt"

func main() {
	x, y := doSum([]int{1, 5, 7, 4, 6, 12, 8}, 55)
	fmt.Println(x, y)
}

//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//你可以按任意顺序返回答案。
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
