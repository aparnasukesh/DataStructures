package main

import "fmt"

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) Append(data int) {
	newNode := &Node{Data: data, Prev: nil, Next: nil}

	if l.Head == nil {
		l.Head, l.Tail = newNode, newNode
		return
	}
	current := l.Tail
	newNode.Prev = current
	current.Next = newNode
	l.Tail = newNode
}

// Access middle element
func (l *LinkedList) MiddleElement() {
	if l.Head == nil {
		fmt.Println("List is empty")
		return
	}
	if l.Head.Next == nil {
		fmt.Println("List has onlyl one element")
		return
	}
	first, second := l.Head, l.Head
	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next.Next
	}
	fmt.Println("Middle element:", first.Data)
}

// Access a data at a specific position
func (l *LinkedList) DataAtSpecificPosition(position int) {
	if position <= 0 {
		fmt.Println("invalid position")
		return
	}
	if l.Head == nil {
		return
	}
	current := l.Head
	for i := 1; i < position && current != nil; i++ {
		current = current.Next
	}
	if current == nil {
		fmt.Println("position is invalid")
		return
	}
	fmt.Printf("position %d contain the data %d", position, current.Data)
}

// Search an element is present or not
func (l *LinkedList) SearchElement(data int) {
	current := l.Head
	for current != nil {
		if current.Data == data {
			fmt.Printf("%d is present", current.Data)
			break
		}
		current = current.Next
	}
	fmt.Println("Data is not present")
}
func (l *LinkedList) ForwardTraversal() {
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
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)
	linkedList.Append(5)

	linkedList.ForwardTraversal()
	fmt.Println("Head:", linkedList.Head.Data, "Tail:", linkedList.Tail.Data)

	linkedList.MiddleElement()
	linkedList.DataAtSpecificPosition(3)
	linkedList.SearchElement(10)
}
