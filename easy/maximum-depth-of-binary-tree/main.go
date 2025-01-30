package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	queue = append(queue, nil)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			depth++
			if len(queue) > 0 {
				queue = append(queue, nil)
			}
		} else {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return depth
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lDepth := maxDepth2(root.Left)
	rDepth := maxDepth2(root.Right)

	if lDepth > rDepth {
		return lDepth + 1
	} else {
		return rDepth + 1
	}
}
