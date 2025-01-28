package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T) {

	p := &TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}
	p.Right = &TreeNode{Val: 3}
	p.Left.Left = &TreeNode{Val: 4}
	p.Left.Right = &TreeNode{Val: 5}
	p.Right.Right = &TreeNode{Val: 6}
	assert.True(t, isSameTree(p, p))

	q := &TreeNode{Val: 1}
	q.Left = &TreeNode{Val: 2}
	q.Right = &TreeNode{Val: 3}
	assert.False(t, isSameTree(p, q))

	assert.False(t, isSameTree(nil, q))
	assert.False(t, isSameTree(p, nil))
}
