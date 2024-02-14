package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

// element added at the end of linkedlist
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

func (l *LinkedList) Traverse() {
	current := l.Head
	for current != nil {
		fmt.Printf("%d ->", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	linkedList := LinkedList{}
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	linkedList.Traverse()
}
