package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func Append(data int, head *ListNode) *ListNode {
	newNode := &ListNode{
		Val:  data,
		Next: nil,
	}
	if head == nil {
		head = newNode
		return head
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
		fmt.Printf("%d->", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	arr := []int{4, 2, 1, 3}
	// Output: [1,2,3,4]
	var head *ListNode
	for i := 0; i < len(arr); i++ {
		head = Append(arr[i], head)
	}

	fmt.Println("\nlist:")
	Traversal(head)

	fmt.Println("\nsorted list:")
	head = sortList(head)
	Traversal(head)
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// step 1: split the list into two halves
	mid := getMiddle(head)
	right := mid.Next
	mid.Next = nil

	// step 2:sort each half recursively
	leftSorted := sortList(head)
	rightSorted := sortList(right)

	// Step 3: Merge the sorted halves
	return merge(leftSorted, rightSorted)
}

func getMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}

	return dummy.Next
}
