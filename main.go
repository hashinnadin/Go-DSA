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

package main

import "fmt"

func main() {
	type Node struct {
		data int
		next *Node
	}

	node1 := Node{data: 10}
	node2 := Node{data: 20}

	node1.next = &node2

	current := &node1

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
