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

func (l *LinkedList) LinkedListToArray() {
	result := []int{}
	current := l.Head
	for current != nil {
		result = append(result, current.Data)
		current = current.Next
	}
	fmt.Println("Array:", result)
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
	linkedList.Append(6)
	linkedList.Append(4)

	linkedList.Traverse()

	linkedList.LinkedListToArray()
}
