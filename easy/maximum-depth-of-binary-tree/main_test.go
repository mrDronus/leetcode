package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {

	r := &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 3}
	r.Left.Left = &TreeNode{Val: 4}
	r.Left.Right = &TreeNode{Val: 5}
	r.Right.Right = &TreeNode{Val: 6}
	r.Right.Right.Right = &TreeNode{Val: 8}
	assert.Equal(t, 4, maxDepth(r))
	assert.Equal(t, 4, maxDepth2(r))

	r = &TreeNode{Val: 1}
	r.Right = &TreeNode{Val: 2}
	assert.Equal(t, 2, maxDepth(r))
	assert.Equal(t, 2, maxDepth2(r))

	assert.Equal(t, 0, maxDepth(nil))
	assert.Equal(t, 0, maxDepth2(nil))
}
