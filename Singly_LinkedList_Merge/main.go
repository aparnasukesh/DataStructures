package main

import (
	"fmt"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

func createLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Value: vals[0], Next: nil}

	current := head

	for _, val := range vals[1:] {
		current.Next = &ListNode{Value: val}
		current = current.Next
	}
	return head
}
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	current1 := list1
	current2 := list2
	var list, tail *ListNode
	for current1 != nil && current2 != nil {
		var newNode *ListNode
		if current1.Value <= current2.Value {
			newNode = &ListNode{Value: current1.Value}
			current1 = current1.Next
		} else {
			newNode = &ListNode{Value: current2.Value}
			current2 = current2.Next
		}

		if list == nil {
			list = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
	}

	for current1 != nil {
		newNode := &ListNode{Value: current1.Value}
		if list == nil {
			list = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
		current1 = current1.Next
	}

	for current2 != nil {
		newNode := &ListNode{Value: current2.Value}
		if list == nil {
			list = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
		current2 = current2.Next
	}
	return list
}
func printLinkedList(head *ListNode) {
	current := head

	for current != nil {
		fmt.Printf("%d ->", current.Value)
		current = current.Next
	}

	fmt.Println("head:", head.Value)

}
func main() {
	list1 := createLinkedList([]int{1, 3, 5})
	list2 := createLinkedList([]int{2, 4, 6})

	printLinkedList(list1)
	fmt.Println()
	printLinkedList(list2)

	mergedList := mergeTwoLists(list1, list2)
	printLinkedList(mergedList)

}
