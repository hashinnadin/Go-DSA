// package main

// import (
// 	"fmt"
// )

// // Node structure
// type Node struct {
// 	data int
// 	next *Node
// }

// // Queue structure
// type Queue struct {
// 	front *Node
// 	rear  *Node
// }

// // Enqueue - Insert at rear
// func (q *Queue) Enqueue(value int) {
// 	newNode := &Node{data: value}

// 	if q.rear == nil {
// 		q.front = newNode
// 		q.rear = newNode
// 		return
// 	}
// 	q.rear.next = newNode
// 	q.rear = newNode
// }

// // Dequeue - Remove from front
// func (q *Queue) Dequeue() int {

// 	if q.front == nil {
// 		fmt.Println("Underflow")
// 	}
// 	value := q.front.data
// 	q.front = q.front.next
// 	return value
// }

// // Peek - View front element
// func (q *Queue) Peek() int {
// 	return q.front.data
// }

// // IsEmpty - Check if queue empty
// func (q *Queue) IsEmpty() bool {
// 	return q.front == nil
// }

// // Print Queue
// func (q *Queue) Print() {
// 	current := q.front
// 	for current != nil {
// 		fmt.Print(current.data, " → ")
// 		current = current.next
// 	}
// 	fmt.Println("nil")
// }

// func main() {

// 	queue := Queue{}

// 	fmt.Println("Enqueue 10, 20, 30")
// 	queue.Enqueue(10)
// 	queue.Enqueue(20)
// 	queue.Enqueue(30)

// 	fmt.Println(queue.front.data, queue.rear.data)

// 	queue.Print()

// 	// fmt.Println("Dequeue:", queue.Dequeue())
// 	queue.Dequeue()
// 	queue.Print()

// 	fmt.Println("Peek:", queue.Peek())

// 	fmt.Println("Is Empty:", queue.IsEmpty())

// 	fmt.Println("Dequeue:", queue.Dequeue())
// 	fmt.Println("Dequeue:", queue.Dequeue())

// 	queue.Print()

// 	fmt.Println("Is Empty:", queue.IsEmpty())
// }

package main

import "fmt"

type Queue struct {
	front *Node
	rear  *Node
}
type Node struct {
	data int
	next *Node
}

func (q *Queue) Enqueue(value int) {
	newNode := &Node{data: value}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
		return
	}
	q.rear.next = newNode
	newNode = q.rear
}

func (Q *Node) Print() {
	current := Q
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
func main() {

	queue := Queue{}

	queue.Enqueue(20)
	queue.front.Print()
}
