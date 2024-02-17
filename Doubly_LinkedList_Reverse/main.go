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

// Reverse
func (l *DoublyLinkedList) Reverse() {
	if l.Head == nil {
		fmt.Println("List is empty")
		return
	}
	if l.Head.Next == nil {
		fmt.Println("List has only one element")
		return
	}

	current := l.Head
	var temp *Node
	for current != nil {
		temp = current.Prev
		current.Prev = current.Next
		current.Next = temp

		current = current.Prev
	}

	if temp != nil {
		l.Head, l.Tail = temp.Prev, temp
	}

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
	//linkedList.BackwardTraversal()
	fmt.Println("Head:", linkedList.Head.Data, "Tail:", linkedList.Tail.Data)

	linkedList.Reverse()
	linkedList.ForwardTraversal()
	fmt.Println("Head:", linkedList.Head.Data, "Tail:", linkedList.Tail.Data)

}
