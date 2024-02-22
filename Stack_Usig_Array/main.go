package main

type Stack struct {
	capacity int
	top      int
	array    []int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		capacity: capacity,
		top:      -1,
		array:    make([]int, capacity),
	}
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) IsFull() bool {
	return s.top == s.capacity-1
}
