package utils

import "errors"

type Stack struct {
	elems  []interface{}
	length int
	top    int
}

func NewStack(length int) *Stack {
	return &Stack{
		elems:  make([]interface{}, length),
		top:    -1,
		length: length,
	}
}

func (s *Stack) Push(elem interface{}) {
	if s.top == s.length-1 {
		s.length = s.length * 2
		var newElems = make([]interface{}, s.length)
		copy(newElems, s.elems)
		s.elems = newElems
	}
	s.top++
	s.elems[s.top] = elem
}

func (s *Stack) Pop() (interface{}, error) {
	if s.top == -1 {
		return 0, errors.New("stack is empty")
	}
	var res = s.elems[s.top]
	s.top--
	return res, nil
}

func (s *Stack) Top() (interface{}, error) {
	if s.top == -1 {
		return 0, errors.New("stack is empty")
	}
	return s.elems[s.top], nil
}

func (s *Stack) Size() int {
	return s.top + 1
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}
