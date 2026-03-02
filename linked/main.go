// package main

// import "fmt"

// type Node struct {
// 	data int
// 	next *Node
// }

// func printList(head *Node) {
// 	current := head

// 	for current != nil {
// 		fmt.Println(current.data)
// 		current = current.next
// 	}
// }

// func insert(head *Node, value int) *Node {
// 	newNode := &Node{data: value}
// 	newNode.next = head
// 	return newNode
// }

// func delete(head *Node) *Node {
// 	if head != nil {
// 		head = head.next
// 	}
// 	return head
// }

// func deleteByValue(head *Node, value int) *Node {

// 	if head.data == value {
// 		return head.next
// 	}
// 	curent := head

// 	for curent.next != nil {
// 		if curent.next.data == value {
// 			curent.next = curent.next.next
// 			break
// 		}
// 		curent = curent.next
// 	}
// 	return head

// }

// func search(head *Node, value int) bool {

// 	current := head
// 	for current.next != nil {
// 		if current.data == value {
// 			return true
// 		}
// 		current = current.next
// 	}
// 	return false
// }

// func searchPosition(head *Node, value int) int {
// 	index := 0
// 	current := head

// 	for current.next != nil {

// 		if current.data == value {
// 			return index
// 		}
// 		current = current.next
// 		index++
// 	}
// 	return -1
// }

// func removeDuplicatesUnsorted(head *Node) *Node {

// 	if head == nil {
// 		return nil
// 	}

// 	seen := make(map[int]bool)
// 	current := head
// 	var prev *Node

// 	for current != nil {

// 		if seen[current.data] {
// 			prev.next = current.next
// 		} else {
// 			seen[current.data] = true
// 			prev = current
// 		}

// 		current = current.next
// 	}

// 	return head
// }

// func main() {

// 	d1 := &Node{data: 10}
// 	d2 := &Node{data: 20}
// 	d3 := &Node{data: 30}
// 	d4 := &Node{data: 40}

// 	d1.next = d2
// 	d2.next = d3
// 	d3.next = d4
// 	head := d1
// 	printList(head)
// 	head1 := search(head, 30)
// 	fmt.Println(head1)
// 	head3 := searchPosition(head, 30)
// 	fmt.Println(head3)
// 	head = deleteByValue(head, 40)
// 	// head = insert(head, 34)
// 	// head = delete(head)
// 	printList(head)
// }

package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func printNode(head *Node) {
	current := head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
func insert(head *Node, value int) *Node {
	newNode := &Node{data: value}
	if head == nil {
		head = newNode
	}
	if head != nil {
		newNode.next = head
		head = newNode
	}
	return head
}

func insertTail(head *Node, value int) *Node {
	newNode := &Node{data: value}

	current := head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	return head
}
func AddvalueByposition(head *Node, value int, pos int) *Node {
	newNode := &Node{data: value}

	if pos == 0 {
		newNode.next = head
		return newNode
	}
	current := head
	for i := 0; current != nil && i < pos-1; i++ {
		current = current.next
	}
	if current != nil {
		newNode.next = current.next
		current.next = newNode
	}
	return head
}
func main() {

	n1 := &Node{data: 1}
	n2 := &Node{data: 2}
	n3 := &Node{data: 3}
	n4 := &Node{data: 4}
	n5 := &Node{data: 5}

	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n5
	head := n1
	printNode(head)
	head = insert(head, 20)
	head = insertTail(head, 80)
	head = AddvalueByposition(head, 88, 3)
	printNode(head)
}
