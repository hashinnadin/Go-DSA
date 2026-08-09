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

// package main

// import "fmt"

// func InsertationSort(arr []int) {

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

// 	arr := []int{10, 2, 4, 34, 90, 4, 3}
// 	InsertationSort(arr)
// 	fmt.Println(arr)
// }

// selection sort //
// package main

// import "fmt"

// func SelectionSort(arr []int) {

// 	n := len(arr)

// 	for i := 0; i < n-1; i++ {
// 		min := i

// 		for j := i + 1; j < n; j++ {
// 			if arr[j] < arr[min] {
// 				min = j
// 			}
// 		}
// 		arr[i], arr[min] = arr[min], arr[i]
// 	}

// }

// func main() {
// 	arr := []int{10, 23, 12, 4, 6, 86}
// 	SelectionSort(arr)
// 	fmt.Println(arr)
// }

// reverce string //

// package main

// import "fmt"

// func Reverce(str string) string {
// 	left := 0
// 	runes := []rune(str)
// 	right := len(runes) - 1

// 	for left < right {
// 		runes[left], runes[right] = runes[right], runes[left]
// 		left++
// 		right--
// 	}
// 	return string(runes)
// }
// func main() {

// 	fmt.Println(Reverce("hashin"))
// }

// package main

// import "fmt"

// func main() {
// 	// var ch rune = 'A'
// 	var str string = "Hashin"
// 	str2 := "you"

// 	fmt.Printf("Hy %s How are %s", str, str2)
// }

// palandrome//

// package main

// import "fmt"

// func CheckPalndrome(str string) bool {

// 	left := 0
// 	right := len(str) - 1

// 	for left < right {
// 		if str[left] != str[right] {
// 			fmt.Println("its not plandrome", str)

// 			return false
// 		}
// 		left++
// 		right--
// 	}
// 	fmt.Println("its plandrome", str)
// 	return true
// }
// func main() {
// 	str := "madam"
// 	fmt.Println(CheckPalndrome(str))
// }

// package main

// import "fmt"

// func Reverce(s string) string {
// 	runes := []rune(s)

// 	left := 0
// 	right := len(runes) - 1

// 	for left < right {
// 		runes[left], runes[right] = runes[right], runes[left]
// 		left++
// 		right--
// 	}
// 	return string(runes)
// }

// func main() {

// 	fmt.Println(Reverce("hashin"))

// }

// Qiuck sort//
package main

import "fmt"

func QiuckSort(arr []int, low, hiegh int) {

	if low < hiegh {
		p := Parti(arr, low, hiegh)
		QiuckSort(arr, low, p-1)
		QiuckSort(arr, p+1, hiegh)
	}
}

func Parti(arr []int, low, hiegh int) int {
	pivot := arr[hiegh]
	i := low - 1

	for j := low; j < hiegh; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[hiegh] = arr[hiegh], arr[i+1]
	return i + 1
}

func main() {

	arr := []int{10, 3, 2, 5, 7}
	QiuckSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// merge sort

// package main

// import "fmt"

// func mergeSort(arr []int) []int {

// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	mid := len(arr) / 2

// 	left := mergeSort(arr[:mid])

// 	right := mergeSort(arr[mid:])

// 	return merge(left, right)

// }

// func merge(left, right []int) []int {
// 	result := make([]int, 0, len(left)+len(right))

// 	i, j := 0, 0

// 	for i < len(left) && j < len(right) {
// 		if left[i] < right[j] {
// 			result = append(result, left[i])
// 			i++
// 		} else {
// 			result = append(result, right[j])
// 			j++
// 		}

// 	}
// 	fmt.Println(left[i:])
// 	result = append(result, left[i:]...)

// 	result = append(result, right[j:]...)
// 	return result

// }

// func main() {

// 	arr := []int{10, 3, 7, 6}

// 	fmt.Println("Before arr :", arr)
// 	fmt.Println(mergeSort(arr))

// }

//heap sort//

// package main

// import "fmt"

// func heapify(arr []int, n, i int) {

// 	largest := i
// 	left := 2*i + 1
// 	right := 2*i + 2

// 	// check the left

// 	if left < n && arr[left] > arr[largest] {
// 		largest = left
// 	}

// 	// check the right

// 	if right < n && arr[right] > arr[largest] {
// 		largest = right
// 	}

// 	// If largest is not the root, swap and continue heapifying
// 	if largest != i {
// 		arr[i], arr[largest] = arr[largest], arr[i]
// 		heapify(arr, n, largest)
// 	}

// }

// func HeapSort(arr []int) {
// 	n := len(arr)

// 	// build max heap
// 	for i := n/2 - 1; i >= 0; i-- {
// 		heapify(arr, n, i)
// 	}

// 	// Exctract element one by one

// 	for i := n - 1; i > 0; i-- {
// 		arr[0], arr[i] = arr[i], arr[0]

// 		heapify(arr, i, 0)
// 	}
// }

// func main() {

// 	arr := []int{4, 10, 5, 7, 9, 10}

// 	fmt.Println("Before Sorting:", arr)

// 	HeapSort(arr)

// 	fmt.Println("After Sorting :", arr)
// }
