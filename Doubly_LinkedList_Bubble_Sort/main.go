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

// Bubble Sort
func (l *LinkedList) BubbleSort() {
	if l.Head == nil || l.Head.Next == nil {
		fmt.Println("list is empty or list has only one value")
		return
	}
	swapped := true
	for swapped {
		swapped = false
		current := l.Head
		for current.Next != nil {
			if current.Data > current.Next.Data {
				current.Data, current.Next.Data = current.Next.Data, current.Data
				swapped = true
			}
			current = current.Next
		}
		if !swapped {
			break
		}

	}
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

	linkedList.Append(6)
	linkedList.Append(1)
	linkedList.Append(4)
	linkedList.Append(9)
	linkedList.Append(3)
	linkedList.Append(5)

	linkedList.ForwardTraversal()
	fmt.Println("Head:", linkedList.Head.Data, "Tail:", linkedList.Tail.Data)

	linkedList.BubbleSort()
	linkedList.ForwardTraversal()
	fmt.Println("Head:", linkedList.Head.Data, "Tail:", linkedList.Tail.Data)

}
