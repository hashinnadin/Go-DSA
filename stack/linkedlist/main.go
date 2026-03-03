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

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

func (S *Stack) Push(value int) {
	newNode := &Node{data: value}
	newNode.next = S.top
	S.top = newNode
}
func (S *Stack) Pop() int {
	value := S.top.data
	S.top = S.top.next
	return value
}

func printList(head *Node) {
	current := head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func deleteHead(head *Node) *Node {
	if head != nil {
		head = head.next
	}
	return head
}

func insertHead(head *Node, value int) *Node {
	newNode := &Node{data: value}
	newNode.next = head
	return newNode
}

func main() {

	head1 := &Node{data: 10}
	head2 := &Node{data: 20}
	head3 := &Node{data: 30}
	head4 := &Node{data: 40}

	head1.next = head2
	head2.next = head3
	head3.next = head4

	head := head1
	head = insertHead(head, 99)
	head = deleteHead(head)
	printList(head)
}
