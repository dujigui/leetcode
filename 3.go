package main

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	l, r, n := 0, 0, 0
	if len(s) != 0 {
		r++
		n++
		m[s[0]] = 1
	}
	for r < len(s) {
		if _, ok := m[s[r]]; ok {
			if n < r-l {
				n = r - l
			}

			for ok {
				delete(m, s[l])
				l++
				_, ok = m[s[r]]
			}
		}
		m[s[r]] = 1
		r++
	}
	if n < r-l {
		n = r - l
	}
	return n
}
