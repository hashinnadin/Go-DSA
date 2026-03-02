// package main

// import "fmt"

// func arrPrint(arr [5]int) {
// 	fmt.Println(arr)
// }

// func main() {
// 	arr := [5]int{1, 2, 3, 4, 5}
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Println(arr[i])
// 	}
// 	arrPrint(arr)
// }

package main

import "fmt"

func reverceArray(arr []int) {
	first := 0
	last := len(arr) - 1

	for first < last {
		arr[first], arr[last] = arr[last], arr[first]
		first++
		last--
	}
}

func duplicates(arr [8]int) []int {

	seen := make(map[int]bool)
	result := []int{}
	for _, value := range arr {
		if !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}
	return result
}

func search(arr []int, value int) bool {
	for _, num := range arr {
		if num == value {
			return true
		}
	}
	return false
}

func BinarySearch(arr [8]int, target int) int {
	low := 0
	hiegh := len(arr) - 1

	for low <= hiegh {
		mir := (low + hiegh) / 2

		if arr[mir] == target {
			return mir
		} else if arr[mir] < target {
			low = mir + 1
		} else {
			hiegh = mir - 1
		}
	}
	return -1
}

func main() {
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	// reverceArray(arr)
	newarr := duplicates(arr)
	found := search(newarr, 30)
	fmt.Println(found)
	binary := BinarySearch(arr, 6)
	fmt.Println(binary)

}
