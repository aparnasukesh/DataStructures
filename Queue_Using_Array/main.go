package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (error, int) {
	if len(q.items) == 0 {
		return fmt.Errorf("Queue is empty"), 0
	}
	item := q.items[0]
	q.items = q.items[1:]
	return nil, item
}

func (q *Queue) Front() (error, int) {
	if len(q.items) == 0 {
		return fmt.Errorf("Queue is empty"), 0
	}
	return nil, q.items[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Traversal() {
	fmt.Print("Queue:")
	for _, item := range q.items {
		fmt.Print(item, " ")
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

	fmt.Println(queue.Size())
	fmt.Println(queue.IsEmpty())
	if err, item := queue.Front(); err == nil {
		fmt.Println("Front:", item)
	}
	queue.Traversal()

	if err, dequeued := queue.Dequeue(); err == nil {
		fmt.Println("Dequeued:", dequeued)
	}
	fmt.Println(queue.Size())
	fmt.Println(queue.IsEmpty())
	if err, item := queue.Front(); err == nil {
		fmt.Println("Front:", item)
	}
	queue.Traversal()

}
