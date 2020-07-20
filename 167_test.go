package main

import (
	"reflect"
	"testing"
)

func Test_twoSum2(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	target := 9
	if !reflect.DeepEqual(twoSum2(numbers, target), []int{1,2}) {
		t.Fail()
	}
}
