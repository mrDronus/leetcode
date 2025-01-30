package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	assert.Equal(t, []int{0, -10, 5, -3, 9}, treeToArray(sortedArrayToBST([]int{-10, -3, 0, 5, 9})))
	assert.Equal(t, []int{1, 3}, treeToArray(sortedArrayToBST([]int{1, 3})))
	assert.Equal(t, []int{2, 1, 3}, treeToArray(sortedArrayToBST([]int{1, 2, 3})))
}

func treeToArray(root *TreeNode) []int {
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
			//r = append(r, node.Val)
			node = node.Right
			continue
		}
		r = append(r, node.Val)
		// leaf of the node, just visit
		if node.Left == nil && node.Right == nil {
			node = nil
			continue
		}

		if node.Left != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return r
}
