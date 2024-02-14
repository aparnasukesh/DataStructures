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

func (l *LinkedList) Traversal() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d ->", current.Data)
		current = current.Next
	}
	fmt.Println(nil)
}

// Delete a node from the beginning
func (l *LinkedList) DeleteFirst() {
	if l.Head == nil {
		fmt.Println("Linked list is empty")
		return
	}
	current := l.Head
	l.Head = l.Head.Next
	current.Next = nil

}

// Delete a Node From the End
func (l *LinkedList) DeleteEnd() {
	if l.Head.Next == nil {
		l.Head = nil
		fmt.Println("Linked list is empty")
		return
	}
	current := l.Head
	for current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil

}

// Delete from specific position
func (l *LinkedList) DeleteAtPosition(position int) {

}

func main() {
	linkedList := &LinkedList{}

	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(5)
	linkedList.Append(4)
	linkedList.Append(8)
	linkedList.Traversal()

	linkedList.DeleteEnd()
	linkedList.Traversal()

	linkedList.DeleteFirst()
	linkedList.Traversal()

	linkedList.DeleteFirst()
	linkedList.Traversal()

	linkedList.DeleteEnd()
	linkedList.Traversal()

	linkedList.DeleteEnd()
	linkedList.Traversal()

}
