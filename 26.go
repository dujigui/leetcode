package main

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums)==1 {
		return len(nums)
	}

	l, r := 1, 1
	for len(nums) > r {
		if nums[r] != nums[l-1] {
			nums[l] = nums[r]
			l++
		}
		r++
	}
	return l
}
