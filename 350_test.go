package main

import (
	"reflect"
	"testing"
)

func TestIntersect1(t *testing.T) {
	n1 := []int{1,2,2,1}
	n2 := []int{2,2}
	r := []int{2,2}
	if !reflect.DeepEqual(intersect(n1, n2), r) {
		t.Fail()
	}
}

func TestIntersect2(t *testing.T) {
	n1 := []int{4,9,5}
	n2 := []int{9,4,9,8,4}
	r := []int{4,9}
	if !reflect.DeepEqual(intersect(n1, n2), r) {
		t.Fail()
	}
}

func TestIntersect3(t *testing.T) {
	n1 := []int{3,1,2}
	n2 := []int{1,1}
	r := []int{1}
	if !reflect.DeepEqual(intersect(n1, n2), r) {
		t.Fail()
	}
}
