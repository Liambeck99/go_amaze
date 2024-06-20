package main

import (
	"errors"
)

type stack []*cell

func NewStack() *stack {
	s := make(stack, 0)
	return &s
}

func (s *stack) Push(p *cell) {
	*s = append(*s, p)
}

func (s *stack) Pop() (*cell, error) {
	l := len(*s)
	if l == 0 {
		return nil, errors.New("empty Stack")
	}

	res := (*s)[l-1]
	*s = (*s)[:l-1]
	return res, nil
}
