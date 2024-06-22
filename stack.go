package main

import (
	"errors"
)

type stack []*cell

func NewStack() *stack {
	s := make(stack, 0)
	return &s
}

func (s *stack) Push(c *cell) {
	*s = append(*s, c)
}

func (s *stack) Pop() (*cell, error) {
	l := len(*s)
	if l == 0 {
		return nil, errors.New("empty stack")
	}

	res := (*s)[l-1]
	*s = (*s)[:l-1]

	return res, nil
}

func (s *stack) Last() (*cell, error) {
	l := len(*s)
	if l == 0 {
		return nil, errors.New("empty stack")
	}

	return (*s)[l-1], nil

}
