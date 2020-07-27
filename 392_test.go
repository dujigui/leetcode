package main

import "testing"

func Test_isSubsequence_1(t *testing.T) {
	s := "abc"
	tt := "ahbgdc"
	if isSubsequence(s, tt) == false {
		t.Fail()
	}
}

func Test_isSubsequence_2(t *testing.T) {
	s := "axc"
	tt := "ahbgdc"
	if isSubsequence(s, tt) == true {
		t.Fail()
	}
}
