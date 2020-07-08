package main

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return []int{}
	}

	if shorter == longer {
		return []int{shorter*k}
	}

	r := make([]int, k+1)
	for i := k; i >= 0; i-- {
		j := k-i
		r[j] = i*shorter+j*longer
	}

	return r
}
