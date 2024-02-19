package main

import (
	"fmt"
)

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

// Remove duplicates
func (l *LinkedList) RemoveDuplicates() {
	exists := make(map[int]bool)
	prev := &Node{}
	current := l.Head

	for current != nil {
		if _, ok := exists[current.Data]; ok {
			prev.Next = current.Next
		} else {
			prev = current
			exists[current.Data] = true
		}
		current = current.Next
	}
}

func (l *LinkedList) Traversal() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d->", current.Data)
		current = current.Next
	}
	fmt.Println(nil)
}

func main() {
	linkedList := LinkedList{}

	linkedList.Append(1)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(3)
	linkedList.Append(3)
	linkedList.Append(5)
	linkedList.Append(1)
	linkedList.Append(4)

	linkedList.Traversal()

	linkedList.RemoveDuplicates()
	linkedList.Traversal()
}
