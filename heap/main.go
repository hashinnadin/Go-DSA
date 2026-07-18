// package main

// import "fmt"

// type MinHeap struct {
// 	data []int
// }

// func (h *MinHeap) Insert(value int) {
// 	h.data = append(h.data, value)

// 	h.heapfyUp(len(h.data) - 1)
// }

// func (h *MinHeap) heapfyUp(index int) {

// 	for index > 0 {

// 		parent := (index - 1) / 2

// 		if h.data[parent] <= h.data[index] {
// 			break
// 		}
// 		h.data[parent], h.data[index] = h.data[index], h.data[parent]
// 		index = parent
// 	}

// }

// func (h *MinHeap) Peek() int {
// 	if len(h.data) == 0 {
// 		return -1
// 	}
// 	return h.data[0]
// }

// func (h *MinHeap) GetMax() int {
// 	max := h.data[0]

// 	for _, value := range h.data {
// 		if value > max {
// 			max = value
// 		}
// 	}
// 	return max
// }

// func (h *MinHeap) Print() {

// 	fmt.Println(h.data)
// }
// func main() {
// 	heap := MinHeap{}

// 	heap.Insert(10)
// 	heap.Insert(20)
// 	heap.Insert(30)
// 	heap.Insert(15)
// 	heap.Insert(5)

// 	heap.Print()
// 	fmt.Println(heap.Peek())
// 	fmt.Println(heap.GetMax())
// }

// package main

// import "fmt"

// func heapify(arr []int, n, i int) {
// 	largest := i
// 	left := 2*i + 1
// 	right := 2*i + 2

// 	if left < n && arr[left] > arr[largest] {
// 		largest = left
// 	}

// 	if right < n && arr[right] > arr[largest] {
// 		largest = right
// 	}

// 	if largest != i {
// 		arr[i], arr[largest] = arr[largest], arr[i]
// 		heapify(arr, n, largest)
// 	}
// }

// func main() {
// 	arr := []int{3, 5, 7, 2, 21, 43, 11, 1, 4}

// 	// Build Max Heap
// 	for i := len(arr)/2 - 1; i >= 0; i-- {
// 		heapify(arr, len(arr), i)
// 	}

// 	fmt.Println(arr)
// }

// package main

// import "fmt"

// type MaxHeap struct {
// 	data []int
// }

// func (h *MaxHeap) Insert(value int) {
// 	h.data = append(h.data, value)
// 	h.heapfyUp(len(h.data) - 1)

// }

// func (h *MaxHeap) heapfyUp(index int) {

// 	for index > 0 {
// 		parent := (index - 1) / 2

// 		if h.data[parent] >= h.data[index] {
// 			break
// 		}
// 		h.data[parent], h.data[index] = h.data[index], h.data[parent]
// 		index = parent
// 	}
// }

// func (h *MaxHeap) Print() {
// 	fmt.Println(h.data)
// }

// func (h *MaxHeap) Peek() int {
// 	if len(h.data) == 0 {
// 		return -1
// 	}
// 	return h.data[0]
// }

// func main() {

// 	heap := MaxHeap{}

// 	heap.Insert(10)
// 	heap.Insert(20)
// 	heap.Insert(30)
// 	heap.Insert(15)
// 	heap.Insert(50)

// 	fmt.Print("Heap: ")
// 	heap.Print()
// 	fmt.Println(heap.Peek())

// }

// // package main

// // import (
// // 	"container/heap"
// // 	"fmt"
// // )

// // // IntHeap implements a Min Heap
// // type IntHeap []int

// // // Number of elements
// // func (h IntHeap) Len() int {
// // 	return len(h)
// // }

// // // Min Heap: smaller element has higher priority
// // func (h IntHeap) Less(i, j int) bool {
// // 	return h[i] < h[j]
// // }

// // // Swap two elements
// // func (h IntHeap) Swap(i, j int) {
// // 	h[i], h[j] = h[j], h[i]
// // }

// // // Push element
// // func (h *IntHeap) Push(x any) {
// // 	*h = append(*h, x.(int))
// // }

// // // Pop element
// // func (h *IntHeap) Pop() any {
// // 	old := *h
// // 	n := len(old)

// // 	item := old[n-1]

// // 	*h = old[:n-1]

// // 	return item
// // }

// // func main() {

// // 	h := &IntHeap{10, 20, 30, 40}

// // 	// Convert slice into a valid heap
// // 	heap.Init(h)

// // 	// Insert
// // 	heap.Push(h, 5)

// // 	fmt.Println("After Push :", *h)

// // 	// Peek
// // 	fmt.Println("Minimum :", (*h)[0])

// // 	// Remove Root
// // 	min := heap.Pop(h)

// // 	fmt.Println("Removed :", min)

// // 	fmt.Println("After Pop :", *h)
// // }

package main

type MaxHeap struct {
	data []int
}

func (h *MaxHeap) insert(value int) {
	h.data = append(h.data, value)
}

func (h *MaxHeap) heapfyUp(index int) {

	for index > 0 {
		parent := (index - 1) / 2
		if h.data[parent] >= h.data[index] {
			break
		}
		h.data[parent], h.h.data[index] = h.data[index], h.data[parent]
		index = parent
	}
}

func main() {

}
