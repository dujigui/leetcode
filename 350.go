package main

// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/

func intersect(nums1 []int, nums2 []int) []int {
	var n1, n2 []int

	// n1 更长
	if len(nums1) > len(nums2) {
		n1 = nums1
		n2 = nums2
	} else {
		n1 = nums2
		n2 = nums1
	}
	m := make(map[int]int, len(n1))
	for _, v := range n1 {
		m[v] += 1
	}

	var r []int
	for _, v := range n2 {
		if e, ok := m[v]; ok && e > 0 {
			r = append(r, v)
			m[v] -= 1
		}
	}
	return r
}
