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

package main

import "fmt"

func Maximum(arr []int) int {
	var result int
	for _, value := range arr {
		if value > 0 {
			result = value
		}
	}
	return result
}

func main() {
	p1 := Maximum([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(p1)
}
