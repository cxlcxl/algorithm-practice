package study_plan

// ListNode 反转链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList 反转链表
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// DeleteDuplicates ...
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// HasCycle 给定一个链表，判断链表中是否有环。
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，
//我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
//如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
//如果链表中存在环，则返回 true 。 否则，返回 false 。
//你能用 O(1)（即，常量）内存解决此问题吗？
//输入：head = [3,2,0,-4], pos = 1
//输出：true
//解释：链表中有一个环，其尾部连接到第二个节点。
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// RemoveElements 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
// 1 2 3 4 5 6
func RemoveElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = RemoveElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

// 合并两个有序链表
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

// SplitListToParts 给你一个头结点为 head 的单链表和一个整数 k ，请你设计一个算法将链表分隔为 k 个连续的部分。
//每部分的长度应该尽可能的相等：任意两部分的长度差距不能超过 1 。这可能会导致有些部分为 null 。
//这 k 个部分应该按照在链表中出现的顺序排列，并且排在前面的部分的长度应该大于或等于排在后面的长度。
//返回一个由上述 k 部分组成的数组。
//输入：head = [1,2,3], k = 5
//输出：[[1],[2],[3],[],[]]
//解释：
//第一个元素 output[0] 为 output[0].val = 1 ，output[0].next = null 。
//最后一个元素 output[4] 为 null ，但它作为 ListNode 的字符串表示是 [] 。
//输入：head = [1,2,3,4,5,6,7,8,9,10], k = 3
//输出：[[1,2,3,4],[5,6,7],[8,9,10]]
//输入被分成了几个连续的部分，并且每部分的长度相差不超过 1 。前面部分的长度大于等于后面部分的长度。
func SplitListToParts(head *ListNode, k int) []*ListNode {
	if head == nil {
		return nil
	}
	cur := head
	listDeep := 1
	for cur.Next != nil {
		listDeep++
		cur = cur.Next
	}
	// 链表公共成员数
	globalMemberNums := listDeep / k
	// 剩余成员
	surplus := listDeep - globalMemberNums*k
	if listDeep < k {
		globalMemberNums = 0
		surplus = listDeep
	}
	listItems := make([]int, k)
	for i := 0; i < k; i++ {
		if surplus > 0 {
			listItems[i] = globalMemberNums + 1
			surplus--
		} else {
			listItems[i] = globalMemberNums
		}
	}
	var retList = make([]*ListNode, k)
	//3 3 2
	cur = head
	for key, item := range listItems {
		retList[key] = cur
		for i := 1; i < item; i++ {
			cur = cur.Next
		}
		cur, cur.Next = cur.Next, nil
	}

	return retList
}
