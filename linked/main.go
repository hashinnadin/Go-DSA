package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func printList(head *Node) {
	current := head

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func insert(head *Node, value int) *Node {
	newNode := &Node{data: value}
	newNode.next = head
	return newNode
}
func main() {

	d1 := &Node{data: 10}
	d2 := &Node{data: 20}
	d3 := &Node{data: 30}
	d4 := &Node{data: 40}

	d1.next = d2
	d2.next = d3
	d3.next = d4

	head := d1
	printList(head)
	head = insert(head, 88)
	printList(head)
}
