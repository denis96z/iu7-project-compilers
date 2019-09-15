package exec

import (
	"github.com/golang-collections/collections/stack"
)

type Stack struct {
	stack *stack.Stack
}

func NewStack() *Stack {
	return &Stack{
		stack: stack.New(),
	}
}

func (s *Stack) Push(object Object) {
	s.stack.Push(object)
}

func (s *Stack) Peek() Object {
	v := s.stack.Peek()
	if v == nil {
		return nil
	}
	return v.(Object)
}

func (s *Stack) Pop() Object {
	v := s.stack.Pop()
	if v == nil {
		return nil
	}
	return v.(Object)
}

func (s *Stack) Count() int {
	return s.stack.Len()
}

func (s *Stack) IsEmpty() bool {
	return s.Count() == 0
}

func (s *Stack) Clear() {
	for !s.IsEmpty() {
		s.Pop()
	}
}
