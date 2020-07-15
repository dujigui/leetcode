package main

import "testing"

func TestIsPalindrome1(t *testing.T) {
	if !isPalindrome(121) {
		t.Fail()
	}
}

func TestIsPalindrome2(t *testing.T) {
	if isPalindrome(-121) {
		t.Fail()
	}
}

func TestIsPalindrome3(t *testing.T) {
	if isPalindrome(10) {
		t.Fail()
	}
}