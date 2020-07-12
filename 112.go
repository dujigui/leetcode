package main

// https://leetcode-cn.com/problems/path-sum/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	return add(root, 0, sum)
}

func add(root *TreeNode, current, sum int) bool {
	if root == nil {
		return false
	}

	addedVal := current + root.Val

	// 叶子
	if root.Left == nil && root.Right == nil {
		return addedVal == sum
	}

	return add(root.Left, addedVal, sum) || add(root.Right, addedVal, sum)
}
