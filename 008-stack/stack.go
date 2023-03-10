package adt

import "errors"

type Stack struct {
	items []string
}

func (s *Stack) Empty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Pop() (string, error) {
	if len(s.items) == 0 {
		return "", errors.New("cannot pop empty stack")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func NewStack() *Stack {
	return &Stack{}
}
