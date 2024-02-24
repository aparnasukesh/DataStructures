package main

import "fmt"

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

func (s *Stack) Push(data int) {
	if s.IsFull() {
		fmt.Println("Stack Overflow")
		return
	}
	s.top++
	s.array[s.top] = data
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		fmt.Println("Stack Underflow")
		return -1
	}
	popped := s.array[s.top]
	s.top--
	return popped
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.array[s.top]
}

func (s *Stack) Traversal() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	for i := s.top; i >= 0; i-- {
		fmt.Println(s.array[i])
	}
}

func main() {
	stack := NewStack(4)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	stack.Push(5)
	fmt.Println(stack.Peek())
	stack.Traversal()
	stack.Pop()
	fmt.Println(stack.Peek())
	stack.Traversal()
}
