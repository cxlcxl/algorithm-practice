package main

type Node struct {
	val   int
	left  *Node
	right *Node
}

func main() {
	nodes := Node{
		val: 100,
		left: &Node{
			val: 75,
			left: &Node{
				val:   70,
				left:  &Node{val: 66},
				right: &Node{val: 73},
			},
			right: &Node{
				val:   80,
				left:  &Node{val: 77},
				right: &Node{val: 90},
			},
		},
		right: &Node{
			val: 110,
			left: &Node{
				val:   100,
				left:  &Node{val: 95},
				right: &Node{val: 105},
			},
			right: &Node{
				val:   120,
				left:  &Node{val: 107},
				right: &Node{val: 125},
			},
		},
	}

	traverse(&nodes)
}

func traverse(root *Node) {
	res := make([]int, 0)
	stack := []*Node{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.val)
		root = root.right
	}
	return

}

// func traverse(root *Node) {
// 	if root == nil {
// 		return
// 	}
// 	traverse(root.left)
// 	fmt.Print(root.val, "\t")
// 	traverse(root.right)
// 	return
// }
