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

func (l *LinkedList) BubbleSort() {
	if l.Head == nil || l.Head.Next == nil {
		fmt.Println("Linked list is empty or it has only one value")
		return
	}
	swapped := true

	for swapped {
		swapped = false
		current := l.Head
		for current.Next != nil {
			if current.Data >= current.Next.Data {
				current.Data, current.Next.Data = current.Next.Data, current.Data
				swapped = true
			}
			current = current.Next
		}
	}
}

func (l *LinkedList) Traversal() {
	current := l.Head

	for current != nil {
		fmt.Printf("%d ->", current.Data)
		current = current.Next
	}
	fmt.Println(nil)
}
func main() {
	linkedList := LinkedList{}
	linkedList.Append(1)
	linkedList.Append(3)
	linkedList.Append(2)
	linkedList.Append(9)
	linkedList.Append(5)
	linkedList.Traversal()

	linkedList.BubbleSort()
	linkedList.Traversal()
}
