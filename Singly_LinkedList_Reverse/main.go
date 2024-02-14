package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Append(data int) {
	newNode := &Node{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Reverse a linked list
func (l *LinkedList) Reverse() {
	var prev *Node
	current, nextNode := l.Head, l.Head

	for nextNode != nil {
		nextNode = current.Next
		current.Next = prev
		prev = current
		current = nextNode

	}
	l.Head = prev
}
func (l *LinkedList) Traversal() {
	current := l.Head

	for current != nil {
		fmt.Printf("%d ->", current.Data)
		current = current.Next
	}
	fmt.Println(nil)
}

func main() {
	linkedList := LinkedList{}

	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)
	linkedList.Traversal()

	linkedList.Reverse()
	linkedList.Traversal()
}
