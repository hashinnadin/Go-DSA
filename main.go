package main

import "fmt"

func main() {

	p1 := 10
	p2 := &p1
	*p2 = 20
	fmt.Println(*p2, p1)
}
