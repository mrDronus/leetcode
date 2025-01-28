package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorderTraversal(t *testing.T) {

	r := &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 3}
	r.Left.Left = &TreeNode{Val: 4}
	r.Left.Right = &TreeNode{Val: 5}
	r.Right.Right = &TreeNode{Val: 6}
	assert.Equal(t, []int{4, 2, 5, 1, 3, 6}, inorderTraversal(r))

	r = &TreeNode{Val: 1}
	r.Right = &TreeNode{Val: 2}
	r.Right.Left = &TreeNode{Val: 3}
	assert.Equal(t, []int{1, 3, 2}, inorderTraversal(r))

}
