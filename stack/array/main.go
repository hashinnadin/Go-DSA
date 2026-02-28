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
func main() {

	var S Stack

	S.Push(10)
	fmt.Println("Hyy", S)
	fmt.Println(S.Pop())
}
