package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type Stack struct {
	top *Node
}

func NewStack() *Stack {
	return &Stack{top: nil}
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Push(data int) {
	newNode := &Node{data: data, next: s.top}
	s.top = newNode
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return -1
	}
	popped := s.top.data
	s.top = s.top.next
	return popped
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.top.data
}

func (s *Stack) Traversal() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	current := s.top
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func main() {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack.Peek())
	stack.Traversal()

	stack.Pop()
	fmt.Println(stack.Peek())
	stack.Traversal()

}
