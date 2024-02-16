package main

import (
	"fmt"
)

type Node struct {
	Data int
	Prev *Node
	Next *Node
}
type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func (l *DoublyLinkedList) Append(data int) {
	newNode := &Node{Data: data, Prev: nil, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	current := l.Tail
	current.Next = newNode
	newNode.Prev = current
	l.Tail = newNode
}

func (l *DoublyLinkedList) ForwardTraversal() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d ->", current.Data)
		current = current.Next
	}
	fmt.Println(nil)

}
func (l *DoublyLinkedList) BackwardTraversal() {
	current := l.Tail

	for current != nil {
		fmt.Printf("%d ->", current.Data)
		current = current.Prev
	}
	fmt.Println(nil)
}
func main() {
	linkedList := DoublyLinkedList{}

	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(5)

	linkedList.ForwardTraversal()
	linkedList.BackwardTraversal()
	fmt.Println("Head:", linkedList.Head.Data, "Tail:", linkedList.Tail.Data)

}
