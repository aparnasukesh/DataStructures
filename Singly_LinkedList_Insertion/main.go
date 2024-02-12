package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Insert(data int) {
	newNode := &Node{data, nil}
	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) Print() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	myLinkedList := LinkedList{}

	myLinkedList.Insert(1)
	myLinkedList.Insert(2)
	myLinkedList.Insert(3)

	fmt.Println("Linked List after Insertions:")
	myLinkedList.Print()
}