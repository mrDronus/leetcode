package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	assert.Equal(t, []int{7, 0, 8}, linkedListToArray(addTwoNumbers(l1, l2)))

	l1 = &ListNode{Val: 0}
	l2 = &ListNode{Val: 0}
	assert.Equal(t, []int{0}, linkedListToArray(addTwoNumbers(l1, l2)))

	l1 = &ListNode{Val: 1}
	l2 = &ListNode{Val: 9}
	assert.Equal(t, []int{0, 1}, linkedListToArray(addTwoNumbers(l1, l2)))

	l1 = &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}
	l2 = &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}
	assert.Equal(t, []int{8, 9, 9, 9, 0, 0, 0, 1}, linkedListToArray(addTwoNumbers(l1, l2)))
}

func linkedListToArray(r *ListNode) []int {
	m := make([]int, 0, 100)
	for r != nil {
		m = append(m, r.Val)
		r = r.Next
	}
	return m
}
