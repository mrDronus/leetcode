package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersRec(l1, l2, false)
}

func addTwoNumbersRec(l1 *ListNode, l2 *ListNode, carry bool) *ListNode {
	if l1 == nil && l2 == nil && !carry {
		return nil
	}
	v := 0
	if carry {
		v++
	}
	var nextL1, nextL2 *ListNode

	if l1 != nil {
		v += l1.Val
		nextL1 = l1.Next
	}
	if l2 != nil {
		v += l2.Val
		nextL2 = l2.Next
	}

	carry = v >= 10
	v = v % 10

	return &ListNode{Val: v, Next: addTwoNumbersRec(nextL1, nextL2, carry)}
}
