package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates1(t *testing.T) {
	nums := []int{1, 1, 2}
	if n := removeDuplicates(nums); n != 2 {
		t.Fail()
	}
	if !reflect.DeepEqual(nums[:2], []int{1, 2}) {
		t.Fail()
	}
}

func TestRemoveDuplicates2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	if n := removeDuplicates(nums); n != 5 {
		t.Fail()
	}
	if !reflect.DeepEqual(nums[:5], []int{0, 1, 2, 3, 4}) {
		t.Fail()
	}
}
