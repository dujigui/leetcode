package main

import "testing"

func TestLengthOfLongestSubstring1(t *testing.T) {
	if lengthOfLongestSubstring("abcabcbb") != 3 {
		t.Fail()
	}
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	if lengthOfLongestSubstring("bbbbb") != 1 {
		t.Fail()
	}
}

func TestLengthOfLongestSubstring3(t *testing.T) {
	if lengthOfLongestSubstring("pwwkew") != 3 {
		t.Fail()
	}
}

func TestLengthOfLongestSubstring4(t *testing.T) {
	if lengthOfLongestSubstring("au") != 2 {
		t.Fail()
	}
}