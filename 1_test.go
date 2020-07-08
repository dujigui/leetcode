package main

import (
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	r := twoSum(nums, target)
	if !reflect.DeepEqual(r, []int{0,1}) {
		t.Fail()
	}
}

func TestTwoSum2(t *testing.T) {
	nums := []int{3,2,4}
	target := 6
	r := twoSum(nums, target)
	if !reflect.DeepEqual(r, []int{1,2}) {
		t.Fail()
	}
}
