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

func insrtTail(head *Node, value int) *Node {
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

func deleteHead(head *Node) *Node {
	if head != nil {
		head = head.next
	}
	return head
}
func deleteTail(head *Node) *Node {
	if head == nil {
		return nil
	}

	if head.next == nil {
		return nil
	}

	current := head

	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
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
	head = deleteHead(head)
	head = AddData(head, 10)
	head = deleteTail(head)
	head = insrtTail(head, 99)
	printData(head)

}
