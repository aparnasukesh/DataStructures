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

// Delete From beginning
func (l *DoublyLinkedList) DeleteFirst() {
	if l.Head == nil {
		fmt.Println("List is empty")
		return
	}
	if l.Head.Next == nil {
		l.Head = nil
		return
	}
	current := l.Head
	l.Head = current.Next
	current.Next = nil
	l.Head.Prev = nil
	// list.Head = list.Head.Next
	// list.Head.Prev = nil

}

// Delete from the end
func (l *DoublyLinkedList) DeleteEnd() {
	if l.Head == nil {
		fmt.Println("List is empty")
		return
	}
	if l.Head.Next == nil {
		l.Head = nil
		return
	}
	l.Tail = l.Tail.Prev
	l.Tail.Next = nil

}

// Delete at specific position
func (l *DoublyLinkedList) DeleteSpecificPosition(position int) {
	if position <= 0 {
		fmt.Println("Invalid position")
		return
	}
	if l.Head.Next == nil {
		l.Head = nil
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
	if current.Next == nil {
		current.Prev.Next = nil
		l.Tail = current.Prev
		current.Prev = nil
		return

	}

	current.Prev.Next = current.Next
	current.Next.Prev = current
	current.Next = nil
	current.Prev = nil

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

	// linkedList.DeleteFirst()
	// linkedList.ForwardTraversal()

	// linkedList.DeleteEnd()
	// linkedList.ForwardTraversal()

	linkedList.DeleteSpecificPosition(2)
	linkedList.ForwardTraversal()

	fmt.Println("Head:", linkedList.Head.Data, "Tail:", linkedList.Tail.Data)

}
