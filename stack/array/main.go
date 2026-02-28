package main

import "fmt"

type Stack struct {
	items []int
}

func (S *Stack) Push(value int) {
	S.items = append(S.items, value)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is Empty")
		return -1
	}
	topindex := len(s.items) - 1
	element := s.items[topindex]
	s.items = s.items[:topindex]
	return element

}

func (s *Stack) Peek() int {
	return s.items[len(s.items)-1]
}

func main() {

	var S Stack

	S.Push(10)
	S.Push(20)
	S.Push(30)
	S.Pop()
	fmt.Println("Hyy", S)
	fmt.Println(S.Peek())
}
