package main

import (
	"testing"
)

func TestAddTwoNumbers1(t *testing.T) {
	l1 := cln(2, cln(4, cln(3, nil)))
	l2 := cln(5, cln(6, cln(4, nil)))

	l3 := addTwoNumbers(l1, l2)

	if l3 == nil || l3.Val != 7 {
		t.Fail()
		return
	}

	l3 = l3.Next
	if l3 == nil || l3.Val != 0 {
		t.Fail()
		return
	}

	l3 = l3.Next
	if l3 == nil || l3.Val != 8 {
		t.Fail()
		return
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	l1 := cln(1, nil)
	ll := l1
	for i := 0; i < 29; i++ {
		ll.Next = cln(0, nil)
		ll = ll.Next
	}
	ll.Next = cln(1, nil)

	l2 := cln(5, cln(6, cln(4, nil)))

	l3 := addTwoNumbers(l1, l2)

	if l3 == nil || l3.Val != 6 {
		t.Fail()
		return
	}

	l3 = l3.Next
	if l3 == nil || l3.Val != 6 {
		t.Fail()
		return
	}

	l3 = l3.Next
	if l3 == nil || l3.Val != 4 {
		t.Fail()
		return
	}

	for i := 0; i < 29-2; i++ {
		l3 = l3.Next
		if l3 == nil || l3.Val != 0 {
			t.Fail()
			return
		}
	}

	l3 = l3.Next
	if l3 == nil || l3.Val != 1 {
		t.Fail()
		return
	}
}

func TestToNumber(t *testing.T) {
	l := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	if 342 != toSmallNumber(l) {
		t.Fail()
	}
}

func TestToList(t *testing.T) {
	l := toSmallList(342)

	if l == nil || l.Val != 2 {
		t.Fail()
		return
	}

	l = l.Next
	if l == nil || l.Val != 4 {
		t.Fail()
		return
	}

	l = l.Next
	if l == nil || l.Val != 3 {
		t.Fail()
		return
	}
}

// create list node
func cln(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
	}
}
