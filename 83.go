package main

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	h := head
	val := head.Val
	for h.Next != nil {
		if h.Next.Val == val {
			h.Next = h.Next.Next
		} else {
			h = h.Next
			val = h.Val
		}
	}
	return head
}
