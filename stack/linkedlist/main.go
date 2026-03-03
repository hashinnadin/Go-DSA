// package main

// import "fmt"

// type Node struct {
// 	data int
// 	next *Node
// }

// type Stack struct {
// 	top *Node
// }

// func (s *Stack) Push(value int) {
// 	newNode := &Node{data: value}
// 	newNode.next = s.top
// 	s.top = newNode
// }
// func (s *Stack) Pop() int {
// 	if s.top == nil {
// 		fmt.Println("")
// 	}
// 	value := s.top.data
// 	s.top = s.top.next
// 	return value
// }

// func (s Stack) Peek() int {
// 	if s.top == nil {
// 		fmt.Println("No Value in stack")
// 		return -1
// 	}
// 	return s.top.data
// }

// func (s *Stack) Print() {
// 	current := s.top

// 	for current != nil {
// 		fmt.Println(current.data)
// 		current = current.next
// 	}
// }
// func main() {
// 	S := Stack{}
// 	S.Push(10)
// 	S.Push(20)
// 	S.Push(30)
// 	S.Push(40)
// 	S.Print()
// 	fmt.Println("------------")
// 	S.Pop()
// 	S.Print()
// 	fmt.Println("Peek :", S.Peek())
// }

// package main

// import "fmt"

// type Node struct {
// 	data int
// 	next *Node
// }

// type Queue struct {
// 	front *Node
// 	rear  *Node
// }

// func (q *Queue) print() {
// 	cuurent := q.front
// 	for cuurent != nil {
// 		fmt.Println(cuurent.data)
// 		cuurent = cuurent.next
// 	}
// }
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

// func (q *Queue) Dequeue() {
// 	if q.front == nil {
// 		fmt.Println("Underflow")
// 	}
// 	q.front = q.front.next
// }

// func (q *Queue) Peek() int {
// 	return q.front.data
// }
// func main() {
// 	Q := Queue{}
// 	Q.Enqueue(10)
// 	Q.Enqueue(20)
// 	Q.Enqueue(30)
// 	Q.Dequeue()
// 	fmt.Println("First value :", Q.Peek())
// 	Q.print()
// }

package main

import "fmt"

func Sum(value, value2 *int) int {
	return *value + *value2
}
func change(x *int) {
	*x = 100
	fmt.Println(*x)
}
func main() {

	a := 10
	b := 20

	result := Sum(&a, &b)
	fmt.Println(result)

	z := 20
	change(&z)

}
