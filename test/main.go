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

// func HeapSort(arr []int) {

// 	n := len(arr)
// 	for i := n - 1; i >= 0; i-- {
// 		heapify(arr, n, i)
// 	}

// 	for i := n - 1; i > 0; i-- {
// 		arr[0], arr[i] = arr[i], arr[0]
// 		heapify(arr, i, 0)
// 	}
// }

// func main() {

// 	arr := []int{10, 2, 4, 3, 8, 1}

// 	HeapSort(arr)
// 	fmt.Println(arr)
// }

package main

import "fmt"

func AddEdge(graph map[string][]string, u, v string) {

	graph[u] = append(graph[u], v)
	graph[v] = append(graph[v], u)
}

func dfs(graph map[string][]string, node string, visited map[string]bool) {

	visited[node] = true
	fmt.Println(node)

	for _, nieg := range graph[node] {
		if !visited[nieg] {
			dfs(graph, nieg, visited)
		}
	}

}

func main() {

	graph := make(map[string][]string)

	visited := make(map[string]bool)

	AddEdge(graph, "A", "B")
	AddEdge(graph, "A", "C")
	AddEdge(graph, "C", "D")
	AddEdge(graph, "D", "E")

	dfs(graph, "A", visited)

}
