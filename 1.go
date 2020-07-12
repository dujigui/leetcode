package main

// https://leetcode-cn.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	r := make([]int, 2)
	m := make(map[int]int)
	for k, v := range nums {
		m[v] = k
	}

	for i, v := range nums {
		j, ok := m[target-v]
		if ok && i != j {
			r[0] = i
			r[1] = j
			break
		}
	}

	return r
}
