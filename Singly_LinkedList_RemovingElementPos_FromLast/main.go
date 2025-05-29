package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 {
		fmt.Println("Invalid input")
		return head
	}

	// Use a dummy node to simplify edge case handling
	dummy := &ListNode{Next: head}
	first := dummy
	second := dummy

	// Move first n+1 steps ahead
	for i := 0; i <= n; i++ {
		if first == nil {
			fmt.Println("Invalid position")
			return head
		}
		first = first.Next
	}

	// Move both first and second until first reaches end
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// Remove the nth node from end
	second.Next = second.Next.Next
	fmt.Println("100 200 300 400 500 ")
	return dummy.Next
}

func Append(head *ListNode, data int) *ListNode {
	newNode := &ListNode{
		Val:  data,
		Next: nil,
	}

	if head == nil {
		return newNode
	}
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	return head
}

func Traversal(head *ListNode) {
	if head == nil {
		fmt.Println("list is empty")
		return
	}

	current := head
	for current != nil {
		fmt.Printf("%d ->", current.Val)
		current = current.Next
	}
	fmt.Printf("nil")
}

func main() {
	var head *ListNode
	head = Append(head, 100)
	Append(head, 200)
	Append(head, 300)
	Append(head, 400)
	Traversal(head)
	fmt.Println("\nHead: ", head.Val)

	removeNthFromEnd(head, 4)
	fmt.Println()
	Traversal(head)
	fmt.Println("\nHead: ", head.Val)
}
