package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// classic recursion
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}
	return isSameTree(p.Right, q.Right)
}
