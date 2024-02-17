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

func ArrayToLinkedList(array []int) *LinkedList {
	linkedList := LinkedList{}
	for _, val := range array {
		linkedList.Append(val)
	}
	return &linkedList
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
	array := []int{1, 2, 3, 4, 5}
	linkedList := ArrayToLinkedList(array)

	linkedList.Traverse()
	fmt.Println("Head:", linkedList.Head.Data)
}
