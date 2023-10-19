package utils

import (
	"calc/tokens"
	"errors"
)

var ErrEmptyStack = errors.New("stack is empty, no elements to pop or retrieve")

type Stack[T any] struct {
	stack []T
}

func (s *Stack[T]) Push(item T) {
	s.stack = append(s.stack, item)
}

func (s *Stack[T]) Top() (T, bool) {
	var nullValue T
	if s.Empty() {
		return nullValue, false
	}
	return s.stack[s.Length()-1], true
}

func (s *Stack[T]) Pop() (T, error) {
	poppedElement, ok := s.Top()
	if !ok {
		return poppedElement, ErrEmptyStack
	}
	s.stack = s.stack[:s.Length()-1]
	return poppedElement, nil
}

func (s *Stack[T]) Length() int {
	return len(s.stack)
}

func (s *Stack[T]) Empty() bool {
	return s.Length() == 0
}

var operators = map[string]struct{}{
	tokens.Add:  {},
	tokens.Div:  {},
	tokens.Subs: {},
	tokens.Mult: {},
	tokens.Lpar: {},
	tokens.Rpar: {},
}

func IsOperator(input string) bool {
	_, ok := operators[input]
	return ok
}
