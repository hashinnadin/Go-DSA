// Bubble Sort//

// package main

// import "fmt"

// func BubbleSort(arr []int) {
// 	n := len(arr)

// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			if arr[j] > arr[j+1] {
// 				arr[j], arr[j+1] = arr[j+1], arr[j]
// 			}
// 		}
// 	}

// }

// func main() {
// 	arr := []int{3, 6, 2, 1, 5, 3, 8}
// 	BubbleSort(arr)
// 	fmt.Println(arr)
// }

//insertaton Sort//

package main

import "fmt"

func InsertationSort(arr []int) {

	for i := 1; i < len(arr); i++ {

		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {

	arr := []int{10, 2, 4, 34, 90, 4, 3}
	InsertationSort(arr)
	fmt.Println(arr)
}
