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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	l1 := list1
	l2 := list2
	var r, current *ListNode
	for {
		if l1 == nil && l2 == nil {
			break
		}

		if current == nil {
			current = &ListNode{}
			r = current
		} else {
			newCurrent := &ListNode{}
			current.Next = newCurrent
			current = newCurrent
		}

		if l1 == nil {
			current.Val = l2.Val
			current.Next = l2.Next
			break
		}
		if l2 == nil {
			current.Next = l1.Next
			current.Val = l1.Val
			break
		}

		if l1.Val <= l2.Val {
			current.Val = l1.Val
			l1 = l1.Next
		} else {
			current.Val = l2.Val
			l2 = l2.Next
		}
	}

	return r
}
