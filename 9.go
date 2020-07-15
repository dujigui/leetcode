package main

// https://leetcode-cn.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	a := make([]int,0)
	for x > 0 {
		a = append(a, x%10)
		x /= 10
	}

	l := len(a)
	for i := 0; i < l/2; i++ {
		if a[i] != a[l-i-1] {
			return false
		}
	}

	return true
}
