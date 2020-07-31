package main

// https://leetcode-cn.com/problems/magic-index-lcci/

func findMagicIndex(nums []int) int {
	for k, v := range nums {
		if k == v {
			return k
		}
	}
	return -1
}
