package main

import (
	studyPlan "algorithm-practice/study-plan"
	"fmt"
)

func main() {
	//fmt.Println(examples.ArrayUnique([]int{0, 0, 0, 1, 1, 2, 3, 3}))
	//fmt.Println(examples.AddNums([]int{2, 4, 3}, []int{5, 6, 4}))
	//fmt.Println(examples.LongestCommonPrefix([]string{"flower", "flow", "fight", ""}))

	//profits := []int{186, 0, 236, 8, 392, 91, 101, 64, 182, 370, 22, 80, 54, 424, 412, 88, 429, 374, 318, 477, 458, 281, 504, 43, 448, 385, 313, 395, 158, 363, 480, 314, 6, 395, 152, 394, 420, 57, 142, 179, 385, 298, 346, 172, 487, 427, 272, 60, 310, 394, 359, 42, 419, 342, 140, 502, 261, 281, 424, 433, 18, 473, 288, 50, 127, 87, 67, 369, 407, 101, 58, 455, 337, 79, 491, 451, 362, 72, 313, 85, 480, 229, 109, 310, 394, 419, 182, 54, 189, 261, 91, 300, 148, 444, 228, 215, 257, 47, 446, 340}
	//capital := []int{75, 486, 155, 104, 72, 136, 174, 194, 368, 121, 258, 445, 160, 383, 73, 18, 437, 308, 509, 482, 227, 469, 104, 416, 257, 97, 88, 82, 181, 169, 463, 56, 182, 249, 467, 140, 328, 291, 115, 339, 511, 73, 53, 373, 220, 261, 236, 296, 284, 431, 178, 94, 520, 196, 150, 172, 26, 487, 96, 285, 433, 404, 204, 130, 313, 374, 89, 140, 401, 261, 76, 370, 126, 230, 73, 509, 377, 446, 480, 504, 61, 82, 504, 85, 241, 17, 84, 412, 18, 174, 469, 10, 449, 114, 215, 340, 414, 82, 401, 61}
	//fmt.Println(examples.FindMaximizedCapital(97, 352, profits, capital))

	// 算法数据结构入门 - 两周计划
	fmt.Println(studyPlan.MaxProfit([]int{2, 8, 1, 5, 3, 9, 4}))
}
