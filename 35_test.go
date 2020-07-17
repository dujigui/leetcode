package main

import "testing"

func Test_searchInsert_1(t *testing.T) {
	if 2 != searchInsert([]int{1,3,5,6},5) {
		t.Fail()
	}
}

func Test_searchInsert_2(t *testing.T) {
	if 1 != searchInsert([]int{1,3,5,6},2) {
		t.Fail()
	}
}

func Test_searchInsert_3(t *testing.T) {
	if 4 != searchInsert([]int{1,3,5,6},7) {
		t.Fail()
	}
}

func Test_searchInsert_4(t *testing.T) {
	if 0 != searchInsert([]int{1,3,5,6},0) {
		t.Fail()
	}
}