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

// Access Middle element
func (l *LinkedList) MiddleElement() {
	first, second := l.Head, l.Head
	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next.Next
	}
	fmt.Println("Middle element is :", first.Data)
}

// Access the data at a specific position
func (l *LinkedList) SpecificPosition(position int) {
	if l.Head == nil || position <= 0 {
		fmt.Println("Invalid position")
		return
	}
	current := l.Head
	for i := 1; i < position && current != nil; i++ {
		current = current.Next
	}
	if current == nil {
		fmt.Println("Position is invalid")
		return
	}
	fmt.Printf("Data at Position %d is %d", position, current.Data)
}

// Search an elemnet is present or not
func (l *LinkedList) SearchElement(data int) {
	current := l.Head
	flag := 0
	for current != nil {
		if current.Data == data {
			flag = 1
			break
		}
		current = current.Next
	}
	if flag == 0 {
		fmt.Println("Data not present")
	} else {
		fmt.Println("Data found")
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

	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(5)
	linkedList.Append(4)
	linkedList.Append(8)
	linkedList.Traversal()

	linkedList.MiddleElement()
	linkedList.SpecificPosition(5)
	linkedList.SearchElement(2)
}
