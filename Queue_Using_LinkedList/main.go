package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) Enqueue(data int) {
	newNode := &Node{data: data, next: nil}
	if q.front == nil {
		q.front = newNode
		q.rear = newNode
		return
	}
	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue) Dequeue() (int, error) {
	if q.front == nil {
		return 0, fmt.Errorf("Queue is empty")
	}
	dequeued := q.front.data
	q.front = q.front.next
	return dequeued, nil

}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

func (q *Queue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("Queue is empty")
	}
	return q.front.data, nil
}

func (q *Queue) Traversal() {
	if q.IsEmpty() {
		return
	}
	current := q.front
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}
func main() {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

}
