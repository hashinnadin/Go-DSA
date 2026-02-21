// package main

// import "fmt"

// func main() {

// 	p1 := 10
// 	p2 := &p1
// 	*p2 = 20
// 	fmt.Println(*p2, p1)
// }

package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	p1 := new(student)
	p1.name = "hashin"
	fmt.Println(p1)

}
