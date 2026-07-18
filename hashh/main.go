// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Frequency(s string) {

// 	words := strings.Fields(s)

// 	freq := make(map[string]int)

// 	for _, v := range words {
// 		freq[v]++
// 	}

// 	for k, v := range freq {
// 		fmt.Println(k, ":", v)
// 	}

// }

// func main() {
// 	str := " go is not easy but go is simple"
// 	Frequency(str)
// // }

// package main

// import "fmt"

// func remove(arr []int) {
// 	check := make(map[int]bool)

// 	for _, v := range arr {
// 		check[v] = true
// 	}

// 	for v := range check {
// 		fmt.Println(v)
// 	}
// }

// func main() {

// 	arr := []int{10, 2, 12, 10, 20, 40, 20, 30, 30}

// 	remove(arr)
// }

// package main

// import "fmt"

// func TwoSum(nums []int, target, target2 int) ([]int, []int) {
// 	m := make(map[int]int)
// 	var res []int
// 	var res2 []int

// 	for i, num := range nums {

// 		need := target - num

// 		if index, ok := m[need]; ok && res2 == nil {
// 			res2 = []int{index, i}
// 		}

// 		need2 := target2 - num

// 		if index, ok := m[need2]; ok && res == nil {
// 			res = []int{index, i}
// 		}
// 		m[num] = i

// 	}
// 	return res, res2
// }

// func main() {
// 	arr := []int{10, 20, 10, 30, 50}
// 	res1, res2 := TwoSum(arr, 40, 50)
// 	fmt.Println(res1)
// 	fmt.Println(res2)

// // }

// package main

// import "fmt"

// func isAnagram(s, t string) bool {
// 	m := make(map[rune]int)

// 	for _, v := range s {
// 		m[v]++
// 	}

// 	for _, v := range t {
// 		m[v]--
// 	}

// 	for _, count := range m {
// 		if count != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	fmt.Println(isAnagram("anagram", "nagaram"))
// 	fmt.Println(isAnagram("peeee", "nagaram"))

// }

package main

import "fmt"

func Duplicates(arr []int) bool {
	m := make(map[int]bool)

	for _, v := range arr {
		if m[v] {
			fmt.Println(v)
			return true
		}
		m[v] = true

	}
	return false

}
func main() {
	arr := []int{10, 20, 30, 10}
	fmt.Println(Duplicates(arr))
}
