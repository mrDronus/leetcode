package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	r := mergeTwoLists(l1, l2)
	m := make([]int, 0, 100)
	for r != nil {
		m = append(m, r.Val)
		r = r.Next
	}
	assert.Equal(t, []int{1, 1, 2, 3, 4, 4}, m)
}
