package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func AddData(head *Node, value int) *Node {
	newNode := &Node{data: value}

	if head == nil {
		return newNode
	}
	current := head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	return head
}
func printData(head *Node) {
	current := head

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
func main() {

	h1 := &Node{data: 10}
	head := h1
	head = AddData(head, 20)
	printData(head)

}
