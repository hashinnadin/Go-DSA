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

package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7}
	// index := 2
	// arr = append(arr[:index], append([]int{99}, arr[index:]...)...)
	arr = arr[1:]
	fmt.Println(arr)

}
