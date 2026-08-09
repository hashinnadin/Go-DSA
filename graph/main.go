// package main

// import "fmt"

// func dfs(graph map[int][]int, node int, visited map[int]bool) {

// 	visited[node] = true

// 	fmt.Print(node, " ")

// 	for _, neighbour := range graph[node] {

// 		if !visited[neighbour] {
// 			dfs(graph, neighbour, visited)
// 		}
// 	}
// }

// func main() {

// 	graph := map[int][]int{
// 		0: {1, 2},
// 		// 1: {0, 3},
// 		// 2: {0, 4},
// 		// 3: {1},
// 		// 4: {2},
// 	}

// 	visited := make(map[int]bool)

// 	dfs(graph, 0, visited)
// }

//bfs

// package main

// import "fmt"

// func bfs(graph map[int][]int, start int) {

// 	visited := make(map[int]bool)

// 	queue := []int{start}
// 	visited[start] = true

// 	for len(queue) > 0 {

// 		current := queue[0]
// 		queue = queue[1:]

// 		fmt.Print(current, " ")

// 		for _, neighbour := range graph[current] {

// 			if !visited[neighbour] {
// 				visited[neighbour] = true
// 				queue = append(queue, neighbour)
// 			}
// 		}
// 	}
// }

// func main() {

// 	graph := map[int][]int{
// 		0: {1, 2},
// 		1: {0, 3},
// 		2: {0, 4},
// 		3: {1},
// 		4: {2},
// 	}

// 	bfs(graph, 0)
// }

// package main

// import "fmt"

// func addEdge(graph map[string][]string, u, v string) {

// 	graph[u] = append(graph[u], v)
// 	graph[u] = append(graph[u], v)
// }

// func dfs(graph map[string][]string, satrt string, visited map[string]bool) {

// 	visited[satrt] = true

// 	fmt.Println(satrt)

// 	for _, neighbour := range graph[satrt] {
// 		if !visited[neighbour] {
// 			dfs(graph, neighbour, visited)

// 		}
// 	}
// }

// func main() {

// 	graph := make(map[string][]string)

// 	visited := make(map[string]bool)

// 	addEdge(graph, "A", "B")
// 	addEdge(graph, "B", "C")
// 	addEdge(graph, "A", "D")
// 	addEdge(graph, "C", "E")

// 	dfs(graph, "A", visited)
// }

package main

import "fmt"

func addEdge(graph map[string][]string, u, v string) {
	graph[u] = append(graph[u], v)
}

func dfs(graph map[string][]string, node string, visited map[string]bool) {

	visited[node] = true

	fmt.Println(node)

	for _, niebour := range graph[node] {
		if !visited[niebour] {
			dfs(graph, niebour, visited)
		}
	}

}

func bfs(graph map[string][]string, start string) {

	visited := make(map[string]bool)

	queue := []string{start}

	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node)

		for _, niebour := range graph[node] {
			if !visited[niebour] {
				visited[niebour] = true
				queue = append(queue, niebour)
			}
		}
	}

}

func main() {

	graph := make(map[string][]string)

	addEdge(graph, "A", "B")
	addEdge(graph, "A", "C")
	addEdge(graph, "B", "D")
	addEdge(graph, "B", "E")
	addEdge(graph, "D", "F")

	visited := make(map[string]bool)
	bfs(graph, "A")
	dfs(graph, "A", visited)

}
