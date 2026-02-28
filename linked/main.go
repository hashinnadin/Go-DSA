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

func delete(head *Node) *Node {
	if head != nil {
		head = head.next
	}
	return head
}

func deleteByValue(head *Node, value int) *Node {
	if head == nil {
		return nil
	}
	if head.data == value {
		head = head.next
	}
	current := head

	for current.next != nil {
		if current.next.data == value {
			current.next = current.next.next
			break
		}
		current = current.next
	}
	return head
}

func search(head *Node, value int) bool {

	current := head

	for current != nil {

		if current.data == value {
			return true
		}
		current = current.next
	}
	return false
}

func searchPosition(head *Node, value int) int {
	index := 0
	current := head

	for current != nil {
		if current.data == value {
			return index
		}
		current = current.next
		index++
	}
	return -1
}

func removeDuplicatesUnsorted(head *Node) *Node {

	if head == nil {
		return nil
	}

	seen := make(map[int]bool)
	current := head
	var prev *Node

	for current != nil {

		if seen[current.data] {
			prev.next = current.next
		} else {
			seen[current.data] = true
			prev = current
		}

		current = current.next
	}

	return head
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
	head1 := search(head, 30)
	fmt.Println(head1)
	head3 := searchPosition(head, 30)
	fmt.Println(head3)
	head = deleteByValue(head, 40)
	head = insert(head, 34)
	// head = delete(head)
	printList(head)
}
