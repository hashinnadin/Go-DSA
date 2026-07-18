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

// package main

// import "fmt"

// func QiuckSort(arr []int, low, high int) {

// 	if low < high {

// 		p := Partition(arr, low, high)

// 		QiuckSort(arr, low, p-1)
// 		QiuckSort(arr, p+1, high)
// 	}
// }

//{10, 4, 8, 2, 6}
//{4, 2, 6, 10, 8} =2
//{2, 4, 6, 10, 8} =0

// func Partition(arr []int, low, high int) int {

// 	pivot := arr[high]
// 	i := low - 1

// 	for j := low; j < high; j++ {
// 		if arr[j] < pivot {
// 			i++
// 			arr[i], arr[j] = arr[j], arr[i]
// 		}
// 	}
// 	arr[i+1], arr[high] = arr[high], arr[i+1]
// 	return i + 1

// }

// func main() {
// 	arr := []int{10, 4, 8, 2, 6}
// 	QiuckSort(arr, 0, len(arr)-1)
// 	fmt.Println(arr)
// }

// merge sort

package main

import "fmt"

func mergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSort(arr[:mid])

	right := mergeSort(arr[mid:])

	return merge(left, right)

}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}

	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result

}

func main() {

	arr := []int{10, 3, 7, 6, 20, 2, 9, 5}

	fmt.Println("Before arr :", arr)
	fmt.Println(mergeSort(arr))

}
