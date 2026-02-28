package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Push(value int) {
	newNode := &Node{data: value}
	newNode.next = s.top
	s.top = newNode
}
func (s *Stack) Pop() int {
	if s.top == nil {
		fmt.Println("")
	}
	value := s.top.data
	s.top = s.top.next
	return value
}

func (s Stack) Peek() int {
	if s.top == nil {
		fmt.Println("No Value in stack")
		return -1
	}
	return s.top.data
}

func (s *Stack) Print() {
	current := s.top

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
func main() {
	S := Stack{}
	S.Push(10)
	S.Push(20)
	S.Push(30)
	S.Push(40)
	S.Print()
	fmt.Println("------------")
	S.Pop()
	S.Print()
	fmt.Println("Peek :", S.Peek())
}
