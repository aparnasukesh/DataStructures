package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

type QueueUsingStack struct {
	Stack1 Stack
	Stack2 Stack
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		fmt.Println("stack is empty")
		return 0, fmt.Errorf("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item, nil

}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (q *QueueUsingStack) Enqueue(data int) {
	q.Stack1.Push(data)
}

func (q *QueueUsingStack) Dequeue() (int, error) {
	if q.Stack2.IsEmpty() {
		for !q.Stack1.IsEmpty() {
			item, _ := q.Stack1.Pop()
			q.Stack2.Push(item)
		}
	}

	return q.Stack2.Pop()
}

func (q *QueueUsingStack) IsEmpty() bool {
	return q.Stack1.IsEmpty() && q.Stack2.IsEmpty()
}

func (q *QueueUsingStack) Size() int {
	return q.Stack1.Size() + q.Stack2.Size()
}

func (q *QueueUsingStack) Traversal() {
	// Traverse stack2 in reverse order
	fmt.Println("Queue:")
	for i := len(q.Stack2.items) - 1; i >= 0; i-- {
		fmt.Printf("%d ", q.Stack2.items[i])
	}

	// Traverse Stack1
	fmt.Println("Queue in Order:")
	for _, iten := range q.Stack1.items {
		fmt.Printf("%d ", iten)
	}
}

func main() {
	// create new queue using stack
	queue := &QueueUsingStack{}

	// enqueue some elements
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Enqueue(40)
	queue.Enqueue(50)
	queue.Enqueue(60)

	queue.Traversal()
	fmt.Println("\nqueue size :", queue.Size())

	if item, err := queue.Dequeue(); err == nil {
		fmt.Println("\nDequeued item: ", item)
	}

	queue.Traversal()

	fmt.Println("is empty?", queue.IsEmpty())
	fmt.Println("\nqueue size :", queue.Size())

	queue.Enqueue(100)
	queue.Enqueue(200)

	queue.Traversal()
	queue.Size()
	fmt.Println("\nqueue size :", queue.Size())

	if item, err := queue.Dequeue(); err == nil {
		fmt.Println("\nDequeued item: ", item)
	}

	queue.Traversal()

	fmt.Println("is empty?", queue.IsEmpty())
	fmt.Println("\nqueue size :", queue.Size())

}
