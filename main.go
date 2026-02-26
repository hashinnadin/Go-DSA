// package main

// import "fmt"

// func main() {

// 	p1 := 10
// 	p2 := &p1
// 	*p2 = 20
// 	fmt.Println(*p2, p1)
// }

// package main

// import "fmt"

// type student struct {
// 	name string
// 	age  int
// }

// func main() {
// 	p1 := new(student)
// 	p1.name = "hashin"
// 	// fmt.Println(p1)
// }

// package main

// import "fmt"

// func main() {
// 	m1 := make(chan int)
//     m1[1]="Hashin"
// }

// package main

// import "fmt"

// func Maximum(arr []int) {

// 	fmt.Println(arr[len(arr)-1])
// }

// func main() {
// 	p1 := []int{1, 2, 3, 4, 5, 6}
// 	Maximum(p1)
// }

// package main

// import "fmt"

// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	index := 2
// 	newValue := 99
// 	fmt.Println(arr)
// 	for i := len(arr) - 1; index > i; i-- {
// 		arr[i] = arr[i-1]
// 		fmt.Println(i)
// 	}
// 	arr[index] = newValue
// 	fmt.Println(arr)
// }

// package main

// import "fmt"

// func main() {

// 	arr := []int{1, 2, 3, 4, 5, 6, 7}
// 	// index := 2
// 	// arr = append(arr[:index], append([]int{99}, arr[index:]...)...)
// 	arr = arr[1:]
// 	arr = arr[:len(arr)-1]
// 	fmt.Println(arr)
// }

// package main

// import "fmt"

// func LinearSearch(arr []int, target int) int {
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == target {
// 			return i
// 		}
// 	}
// 	return -1
// }
// n
// func main() {
// 	p1 := LinearSearch([]int{1, 2, 3, 4, 5}, int(7))
// 	fmt.Println(p1)
//

// package main

// import "fmt"

// func BinarySearch(arr []int, target int) int {
// 	low := 0
// 	hiegh := len(arr) - 1

// 	for low <= hiegh {
// 		mid := (low + hiegh) / 2

// 		if arr[mid] == target {
// 			return mid
// 		} else if arr[mid] < target {
// 			low = mid + 1
// 		} else {
// 			hiegh = mid - 1
// 		}
// 	}
// 	return -1
// }

// func main() {
// 	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140}
// 	targer := 30
// 	fmt.Println(BinarySearch(arr, targer))
// }

// package main

// import "fmt"

// func ReverseArray(arr []int) {
// 	left := 0
// 	rieght := len(arr) - 1

// 	for left < rieght {
// 		arr[left], arr[rieght] = arr[rieght], arr[left]
// 		left++
// 		rieght--
// 	}
// }
// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7}
// 	ReverseArray(arr)
// 	fmt.Println(arr)
// }

// package main

// import "fmt"

// func main() {
// 	arr := []int{1, 2, 3, 4, 5}
// 	index := 1
// 	arr = append(arr[:index], append([]int{99}, arr[index:]...)...)
// 	// arr = append(arr[:2], arr[3:]...)
// 	fmt.Println(arr)
// }

// package main

// import "fmt"

// func main() {

// 	arr := []int{1, 2, 3, 4, 5}

// 	low := 0
// 	hiegh := len(arr) - 1

// 	for low < hiegh {

// 		arr[low], arr[hiegh] = arr[hiegh], arr[low]
// 		low++
// 		hiegh--
// 	}
// 	fmt.Println(arr)
// }

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

// func main() {

// 	node1 := &Node{data: 10}
// 	node2 := &Node{data: 20}
// 	node3 := &Node{data: 30}

// 	node1.next = node2
// 	node2.next = node3

// 	head := node1

//  // insert HEAD
// 	newNode := &Node{data: 5}
// 	newNode.next = head
// 	head = newNode

// 	printList(head)

// }
package main

import "fmt"

type Node struct {
	data int
	next *Node
}

// Print List
func printList(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.data, " → ")
		current = current.next
	}
	fmt.Println("nil")
}

func insertAtPosition(head *Node, value int, pos int) *Node {
	newNode := &Node{data: 99}

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

func insertAtend(head *Node, value int) *Node {
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

// Delete by Value
func deleteByValue(head *Node, value int) *Node {

	if head == nil {
		return nil
	}

	// If head must be deleted
	if head.data == value {
		return head.next
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

func main() {

	node1 := &Node{data: 5}
	node2 := &Node{data: 10}
	node3 := &Node{data: 15}
	node4 := &Node{data: 20}
	node5 := &Node{data: 10}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5

	head := node1

	fmt.Println("Original List:")
	printList(head)

	// Insert 100 at position 2
	head = insertAtPosition(head, 100, 2)

	fmt.Println("After Insert 100 at position 2:")
	printList(head)

	// Delete value 10
	head = deleteByValue(head, 10)

	head = insertAtend(head, 100)
	fmt.Println("After Delete value 10:")
	printList(head)
}
