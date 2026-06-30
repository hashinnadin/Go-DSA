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

package main

import "fmt"

func Reverce(s string) string {
	runes := []rune(s)

	left := 0
	right := len(runes) - 1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)
}

func main() {

	fmt.Println(Reverce("hashin"))

}
