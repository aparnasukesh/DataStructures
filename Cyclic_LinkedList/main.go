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

// hasCycle function checks if a linked list has a cycle
func hasCycle(head *Node) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}

// InsertAftreValue inserts a new node with given value after the first node with the specified value in a cyclic linked list
func insertAfterValue(head *Node, target int, value int) {
	if head == nil {
		return
	}

	current := head

	for {

		if current.Data == target {
			newNode := &Node{
				Data: value,
			}
			newNode.Next = current.Next
			current.Next = newNode
			break
		}

		current = current.Next

		if current == head {
			// We've traversed the whole cycle and didn't find the target value
			fmt.Println("Target value not found in the list")
			break
		}
	}
}

// deleteValue deletes the first occurrence of the specified value in a cyclic linked list
func deleteValue(head *Node, target int) *Node {
	if head == nil {
		return nil
	}

	current := head
	var prev *Node

	for {

		if current.Data == target {
			if prev != nil {
				prev = current.Next
			} else {
				// If the target is the head, we need to find the last node
				last := head
				for last.Next != head {
					last = last.Next
				}
				if head == head.Next { // Only one node in the list
					return nil
				}
				head = head.Next
				last.Next = head
			}
			return head
		}
		prev = current
		current = current.Next

		if current == head {
			// We've traversed the whole cycle and didn't find the target value
			fmt.Println("Target value not found in the list")
			break
		}
	}
	return head
}

// Create or Make Cycle
func makeCycle(head *Node, pos int) {
	current := head
	var startNode *Node
	count := 1

	for current.Next != nil {
		if count == pos {
			startNode = current
		}
		current = current.Next
		count++
	}
	current.Next = startNode
}

// RemoveCycle removes the cycle from the list
func removeCycle(head *Node) {
	cycleStart := findCycleStart(head)
	if cycleStart == nil {
		return // no cycle
	}

	// find the node just before the cycle start
	current := cycleStart
	for current.Next != cycleStart {
		current = current.Next
	}

	// break the cycle
	current.Next = nil
}

// FindCycleStart finds the starting node of the cycle in the linked list
func findCycleStart(head *Node) *Node {
	if head == nil {
		return nil
	}
	slow := head
	fast := head

	// Delete if there is a cycle
	hasCycle := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			hasCycle = true
			break
		}
	}

	// if there is no cycle , return nil
	if !hasCycle {
		return nil
	}

	// Find the start of the cycle
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// printList prints the linked list ( for debugging purpose )
func printList(head *Node, limit int) {
	current := head

	count := 0

	for current != nil && count < limit {
		fmt.Printf("%d ->", current.Data)
		current = current.Next
		count++
	}
	fmt.Println(".....")
}

func main() {

	// create the node for linked list
	head := &Node{Data: 10}
	first := &Node{Data: 1}
	second := &Node{Data: 2}
	third := &Node{Data: 3}
	fourth := &Node{Data: 4}
	fifth := &Node{Data: 5}

	// Link the nodes to form a cyclic list :1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 3
	head.Next = first
	first.Next = second
	second.Next = third
	third.Next = fourth
	fourth.Next = fifth
	fifth.Next = nil // Create a cycle here

	// Print the original cyclic linked list (up to a limit to avoid infinite loop)
	fmt.Println("Original cyclic linked list:")
	printList(head, 10)

	// Insert a new value after a specific value in the cyclic linked list
	insertAfterValue(head, 5, 600)

	// Print the updated cyclic linked list (up to a limit to avoid infinite loop)
	fmt.Println("Updated cyclic linked list after inserting 6 after 3:")
	printList(head, 15)

	// Check if the linked list is cyclic or not  (Ckeck if the linked list form a cycle )
	if hasCycle(head) {
		fmt.Println("The linked list has a cycle ")
	} else {
		fmt.Println("The linked list doesn't have a cycle")
	}

	makeCycle(head, 4)
	printList(head, 10)

	removeCycle(head)
	printList(head, 10)

}
