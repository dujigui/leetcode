package main

import (
	"testing"
)

func TestHasPathSum(t *testing.T) {
	r := createRoot()
	if !hasPathSum(r, 22) {
		t.Fail()
	}
}

func createNode(val int, left, right *TreeNode) *TreeNode {
	r := new(TreeNode)
	r.Val = val
	r.Left = left
	r.Right = right
	return r
}

func createRoot() *TreeNode {
	return createNode(5,
		createNode(4,
			createNode(11,
				createNode(7, nil, nil),
				createNode(2, nil, nil)),
			nil),
		createNode(8,
			createNode(13, nil, nil),
			createNode(4,
				nil,
				createNode(1, nil, nil))),
	)
}
