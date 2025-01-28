package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// replace recursion with stack
func inorderTraversal(root *TreeNode) []int {
	r := []int{}
	stack := []*TreeNode{}
	node := root
	for {
		// traversed
		if len(stack) == 0 && node == nil {
			break
		}
		// pop from stack, visit and go right
		if node == nil {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			r = append(r, node.Val)
			node = node.Right
			continue
		}
		// leaf of the node, just visit
		if node.Left == nil && node.Right == nil {
			r = append(r, node.Val)
			node = nil
			continue
		}

		if node.Left != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			r = append(r, node.Val)
			node = node.Right
		}
	}
	return r
}
