package main

import "testing"

func Test_findMagicIndex_1(t *testing.T) {
	if 0 != findMagicIndex([]int{0,2,3,4,5}) {
		t.Fail()
	}
}

func Test_findMagicIndex_2(t *testing.T) {
	if 1 != findMagicIndex([]int{1,1,1}) {
		t.Fail()
	}
}