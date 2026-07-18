// package main

// import "fmt"

// func Insert(arr []int) {

// 	for i := 1; i < len(arr); i++ {

// 		key := arr[i]
// 		j := i - 1

// 		for j >= 0 && arr[j] > key {
// 			arr[j+1] = arr[j]
// 			j--
// 		}
// 		arr[j+1] = key
// 	}
// }

// func main() {

// 	arr := []int{10, 3, 8, 20, 1}
// 	Insert(arr)
// 	fmt.Println(arr)
// }

// package main

// import "fmt"

// func Frequency(a string) {
// 	ferq := make(map[rune]int)

// 	for _, v := range a {
// 		ferq[v]++
// 	}
// 	for k, v := range ferq {
// 		fmt.Println(string(k), v)
// 	}
// }

// func main() {

// 	s := "hashin"
// 	Frequency(s)
// }

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func Insert(root *Node, value int) *Node {

	if root == nil {
		root = &Node{data: value}
		return root
	}

	if root.data > value {
		root.left = Insert(root.left, value)
	} else if root.data < value {
		root.right = Insert(root.right, value)
	}
	return root
}

func Inorder(root *Node) *Node {

	if root == nil {
		return nil
	}
	Inorder(root.left)
	fmt.Println(root.data)
	Inorder(root.right)

	return root
}

func print(root *Node) {
	fmt.Println(root.data)
}

func main() {

	root := &Node{}

	Insert(root, 10)
	Insert(root, 20)
	Insert(root, 40)
	Insert(root, 4)
	Insert(root, 8)

	Inorder(root)
}
