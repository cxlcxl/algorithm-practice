package main

func main() {
	// doSum
	// x, y := doSum([]int{1, 5, 7, 4, 6, 12, 8}, 55)
	// fmt.Println(x, y)
}

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。
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

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
