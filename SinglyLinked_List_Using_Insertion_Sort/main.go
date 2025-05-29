package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Val: -1}
	current := head

	for current != nil {
		prev := dummy
		next := current.Next

		// Find the correct position to insert current
		for prev.Next != nil && prev.Next.Val < current.Val {
			prev = prev.Next
		}

		// Insert current between prev and prev.Next
		current.Next = prev.Next
		prev.Next = current

		// Move to the next node in the original list
		current = next
	}

	return dummy.Next
}

func insertionSortList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	curr := head.Next
	prev := head
	for curr != nil {
		if curr.Val < prev.Val {
			temp := dummy
			for temp.Next.Val < curr.Val {
				temp = temp.Next
			}
			prev.Next = curr.Next
			curr.Next = temp.Next
			temp.Next = curr
			curr = prev.Next
		} else {
			prev = curr
			curr = curr.Next
		}
	}
	return dummy.Next
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
		fmt.Printf("%d ->", current.Val)
		current = current.Next
	}
}
func main() {
	arr := []int{2, 1, 3, 6, 100, 60, 700, 89, 556, 7890, 88, 4, 9, 7, 5}

	// for i := 0; i < len(arr); i++ {
	// 	j := i - 1
	// 	val := arr[i]

	// 	for j >= 0 && arr[j] > val {
	// 		arr[j+1] = arr[j]
	// 		j--
	// 	}
	// 	arr[j+1] = val
	// }
	// fmt.Println(arr)

	var head *ListNode
	for i := 0; i < len(arr); i++ {
		head = Append(arr[i], head)
	}
	Traversal(head)
	fmt.Println("\n", head.Val)

	fmt.Println("\nafter sorting the list")
	// head = insertionSortList(head)
	// Traversal(head)
	// fmt.Println("\n", head.Val)
	head = insertionSortList1(head)
	Traversal(head)
}
