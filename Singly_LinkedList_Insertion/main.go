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

// Insertion at the end of the linked list
func (l *LinkedList) Append(data int) {
	newNode := &Node{
		Data: data,
		Next: nil,
	}

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

// Insertio at the beginning of the linked list
func (l *LinkedList) InsertionAtBegin(data int) {
	newNode := &Node{
		Data: data,
		Next: l.Head,
	}
	l.Head = newNode

}

// Insertion at the specific position of the linked list
func (l *LinkedList) InsertionAtPositon(data, position int) {
	newNode := &Node{
		Data: data,
		Next: nil,
	}
	current := l.Head

	for i := 1; i < position-1; i++ {
		current = current.Next
	}
	if current == nil {
		fmt.Println("Invalid position")
		return
	}
	newNode.Next = current.Next
	current.Next = newNode
}

// Traversal of linked list
func (l *LinkedList) Traversal() {
	current := l.Head

	for current != nil {
		fmt.Printf("%d ->", current.Data)
		current = current.Next
	}
	fmt.Println(nil)

}

// Length of linked list
func (l *LinkedList) Lenghth() {
	count := 0
	current := l.Head
	for current != nil {
		current = current.Next
		count++
	}
	fmt.Println("Length:", count)

}
func main() {
	linkedList := LinkedList{}

	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)

	linkedList.Traversal()

	linkedList.InsertionAtBegin(100)
	linkedList.InsertionAtBegin(200)
	linkedList.Traversal()

	linkedList.InsertionAtPositon(500, 3)
	linkedList.InsertionAtPositon(1000, 5)
	linkedList.Traversal()

	linkedList.Lenghth()

}
