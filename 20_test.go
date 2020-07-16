package main

import "testing"

func Test_isValid_1(t *testing.T) {
	if !isValid("()") {
		t.Fail()
	}
}

func Test_isValid_2(t *testing.T) {
	if !isValid("()[]{}") {
		t.Fail()
	}
}

func Test_isValid_3(t *testing.T) {
	if isValid("(]") {
		t.Fail()
	}
}

func Test_isValid_4(t *testing.T) {
	if isValid("([)]") {
		t.Fail()
	}
}

func Test_isValid_5(t *testing.T) {
	if !isValid("{[]}") {
		t.Fail()
	}
}

func Test_stack_push(t *testing.T) {
	s := stack{bytes: make([]int32, 0, 10)}
	s.push(1)
	s.push(2)
	if len(s.bytes) != 2 {
		t.Fail()
	}
	if s.bytes[0] != 1 {
		t.Fail()
	}
	if s.bytes[1] != 2 {
		t.Fail()
	}
}

func Test_stack_pop(t *testing.T) {
	s := stack{bytes: make([]int32, 0, 10)}
	s.push(1)
	s.push(2)
	if len(s.bytes) != 2 {
		t.Fail()
	}
	if b := s.pop(); b != 2 {
		t.Fail()
	}
	if len(s.bytes) != 1 {
		t.Fail()
	}
}

func Test_stack_peek(t *testing.T) {
	s := stack{bytes: make([]int32, 0, 10)}
	s.push(1)
	s.push(2)
	if len(s.bytes) != 2 {
		t.Fail()
	}
	if b := s.peek(); b != 2 {
		t.Fail()
	}
	if len(s.bytes) != 2 {
		t.Fail()
	}
}
