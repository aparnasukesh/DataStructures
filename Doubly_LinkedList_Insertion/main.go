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

// Insertion At end
func (l *DoublyLinkedList) Append(data int) {
	newNode := &Node{Data: data, Next: nil, Prev: nil}
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

// Insertion at the beginning
func (l *DoublyLinkedList) InsertionAtBeginning(data int) {
	newNode := &Node{Data: data, Prev: nil, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	current := l.Head
	newNode.Next = current
	current.Prev = newNode
	l.Head = newNode

}

// Insertion at a specific position
func (l *DoublyLinkedList) InsertionAtSpecificPosition(data, position int) {
	newNode := &Node{Data: data, Prev: nil, Next: nil}
	if l.Head == nil {
		l.Head, l.Tail = newNode, newNode
		return
	}
	if position <= 0 {
		fmt.Println("Invalid position")
		return
	}
	current := l.Head
	for i := 1; i < position && current != nil; i++ {
		current = current.Next
	}
	if current == nil {
		fmt.Println("Invalid position")
		return
	}
	newNode.Next = current
	newNode.Prev = current.Prev
	current.Prev.Next = newNode
	current.Prev = newNode

}

// Insertion at specific position ,we also do like this
func (list *DoublyLinkedList) InsertAt(data, position int) {
	newNode := &Node{Data: data, Prev: nil, Next: nil}

	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		return
	}

	if position <= 0 {
		newNode.Next = list.Head
		list.Head.Prev = newNode
		list.Head = newNode
		return
	}

	current := list.Head
	for i := 0; i < position-1 && current != nil; i++ {
		current = current.Next
	}

	if current == nil {
		// Insert at the end if the position is beyond the current size of the list
		newNode.Prev = list.Tail
		list.Tail.Next = newNode
		list.Tail = newNode
	} else {
		newNode.Prev = current
		newNode.Next = current.Next
		if current.Next != nil {
			current.Next.Prev = newNode
		}
		current.Next = newNode
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
	linkedList.InsertionAtBeginning(1)
	linkedList.InsertionAtBeginning(2)
	linkedList.InsertionAtBeginning(3)

	linkedList.Append(6)
	linkedList.Append(7)

	linkedList.ForwardTraversal()

	linkedList.InsertionAtSpecificPosition(100, 3)
	linkedList.ForwardTraversal()

	fmt.Println("Head:", linkedList.Head.Data, "Tail:", linkedList.Tail.Data)
}
