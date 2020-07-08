package main

import (
	"reflect"
	"testing"
)

func TestDivingBoard(t *testing.T) {
	r := divingBoard(1, 2, 3)
	if !reflect.DeepEqual(r, []int{3,4,5,6}) {
		t.Fail()
	}
}