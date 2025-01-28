package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T) {

	p := &TreeNode{Val: 1}
	p.Left = &TreeNode{Val: 2}
	p.Left.Left = &TreeNode{Val: 3}
	p.Left.Right = &TreeNode{Val: 4}

	p.Right = &TreeNode{Val: 2}
	p.Right.Left = &TreeNode{Val: 4}
	p.Right.Right = &TreeNode{Val: 3}

	assert.True(t, isSymmetric(p))

	p.Left.Right = nil
	assert.False(t, isSymmetric(p))
}
