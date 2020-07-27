package main

// https://leetcode-cn.com/problems/is-subsequence/

func isSubsequence(s string, t string) bool {
	ls, lt := len(s), len(t)
	is, it := 0, 0
	for is < ls && it < lt {
		if s[is] == t[it] {
			is++
		}
		it++
	}
	return is == ls
}
