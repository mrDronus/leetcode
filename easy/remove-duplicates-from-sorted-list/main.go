package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	r := &ListNode{Val: head.Val}
	resultHead := r
	currentHead := head
	for currentHead != nil {
		if r.Val == currentHead.Val {
			currentHead = currentHead.Next
			continue
		}
		n := &ListNode{Val: currentHead.Val}
		r.Next = n
		r = n
	}
	return resultHead
}
