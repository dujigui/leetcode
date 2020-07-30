package main

import (
	"testing"
)

func Test_deleteDuplicates_1(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	head = deleteDuplicates(head)
	if head.Val != 1 {
		t.Fail()
	}
	if head.Next.Val != 2 {
		t.Fail()
	}
	if head.Next.Next != nil {
		t.Fail()
	}
}

func Test_deleteDuplicates_2(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	head = deleteDuplicates(head)
	if head.Val != 1 {
		t.Fail()
	}
	if head.Next.Val != 2 {
		t.Fail()
	}
	if head.Next.Next.Val != 3 {
		t.Fail()
	}
	if head.Next.Next.Next != nil {
		t.Fail()
	}
}
