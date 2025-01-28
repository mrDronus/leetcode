package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicates(t *testing.T) {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	r := deleteDuplicates(list)
	m := make([]int, 0, 100)
	for r != nil {
		m = append(m, r.Val)
		r = r.Next
	}
	assert.Equal(t, []int{1, 2}, m)

	list = &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}}
	r = deleteDuplicates(list)
	m = make([]int, 0, 100)
	for r != nil {
		m = append(m, r.Val)
		r = r.Next
	}
	assert.Equal(t, []int{1, 2, 3}, m)
}
