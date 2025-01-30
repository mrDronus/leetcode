package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTrec(nums, 0, len(nums)-1)
}

func sortedArrayToBSTrec(nums []int, from, to int) *TreeNode {
	if from > to {
		return nil
	}

	middle := from + (to-from)/2
	r := &TreeNode{Val: nums[middle]}
	r.Left = sortedArrayToBSTrec(nums, from, middle-1)
	r.Right = sortedArrayToBSTrec(nums, middle+1, to)
	return r
}
