package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// with recursion
func isSymmetric(root *TreeNode) bool {
	return isLREqual(root.Left, root.Right)
}

func isLREqual(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	return isLREqual(left.Left, right.Right) && isLREqual(left.Right, right.Left)
}
