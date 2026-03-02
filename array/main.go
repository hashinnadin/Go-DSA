package main

import "fmt"

func arrPrint(arr [5]int) {
	fmt.Println(arr)
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	arrPrint(arr)
}
