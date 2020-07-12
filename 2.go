package main

import (
	"math"
)

// https://leetcode-cn.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		r       *ListNode
		current *ListNode
		n       int
		l1Val   int
		l2Val   int
	)

	if l1 == nil && l2 == nil {
		return &ListNode{}
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	for {
		if l1 == nil {
			l1Val = 0
		} else {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			l2Val = 0
		} else {
			l2Val = l2.Val
			l2 = l2.Next
		}

		n = n + l1Val + l2Val
		l := &ListNode{
			Val:  n % 10,
			Next: nil,
		}
		n = n / 10

		if current != nil {
			current.Next = l
		}
		current = l
		if r == nil {
			r = current
		}

		if l1 == nil && l2 == nil {
			if n != 0 {
				current.Next = &ListNode{
					Val:  n % 10,
					Next: nil,
				}
				n = n / 10
			}
			break
		}
	}

	return r
}

// 只能用于整型所能表示正整数范围内的场景
// return toSmallList(toSmallNumber(l1) + toSmallNumber(l2))
func toSmallNumber(l *ListNode) int {
	n := 0
	r := 0
	for {
		if l == nil {
			break
		}
		r += l.Val * int(math.Pow(10, float64(n)))
		l = l.Next
		n++
	}
	return r
}

// 只能用于整型所能表示正整数范围内的场景
func toSmallList(i int) *ListNode {
	if i == 0 {
		return &ListNode{
			Val:  0,
			Next: nil,
		}
	}

	var r, current *ListNode
	for {
		if i == 0 {
			break
		}

		if current == nil {
			current = &ListNode{}
		} else {
			current.Next = &ListNode{}
			current = current.Next
		}
		if r == nil {
			r = current
		}

		current.Val = i % 10
		i = i / 10
	}

	return r
}
