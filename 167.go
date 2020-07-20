package main

// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/

func twoSum2(numbers []int, target int) []int {
	for index1, v := range numbers {
		if index2 := find(numbers, index1+1, len(numbers)-1, target-v); index2 != -1 {
			return []int{index1+1, index2+1}
		}
	}
	return []int{}
}

func find(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}
	if start == end {
		if nums[start] == target {
			return start
		} else {
			return -1
		}
	}

	mid := (end-start)/2 + start

	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return find(nums, start, mid-1, target)
	} else {
		return find(nums, mid+1, end, target)
	}
}
