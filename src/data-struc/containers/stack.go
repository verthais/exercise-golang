package containers

import (
	"errors"
	"exercise-go/src/data-struc/list"
)

type Stack struct {
	data list.LinkedList
}

func NewStack() Stack {
	return Stack{data: list.NewLinkedList()}
}

func (s *Stack) Push(value interface{}) {
	s.data.AddFront(value)
}

func (s *Stack) Pop() (interface{}, error) {
	v, e := s.data.PopFront()

	if e != nil {
		e = errors.New("Cannot pop from empty Stack")
	}

	return v, e
}

func (s *Stack) Size() int {
	return s.data.Size
}
