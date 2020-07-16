package main

// https://leetcode-cn.com/problems/valid-parentheses/

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}

	ss := stack{bytes: make([]int32, 0, len(s)/2)}
	for _, v := range s {
		if isPair(ss.peek(), v) {
			ss.pop()
		} else {
			ss.push(v)
		}
	}

	return ss.len() == 0
}

func isPair(a, b int32) bool {
	if a == 40 && b == 41 {
		return true
	} else if a == 91 && b == 93 {
		return true
	} else if a == 123 && b == 125 {
		return true
	} else {
		return false
	}
}

// 简化逻辑，不考虑部分边界情况
type stack struct {
	bytes []int32
}

func (s *stack) push(b int32) {
	s.bytes = append(s.bytes, b)
}

func (s *stack) pop() int32 {
	r := s.bytes[len(s.bytes)-1]
	s.bytes = s.bytes[:len(s.bytes)-1]
	return r
}

func (s *stack) peek() int32 {
	if len(s.bytes) == 0 {
		return -1
	}
	return s.bytes[len(s.bytes)-1]
}

func (s *stack) len() int {
	return len(s.bytes)
}
