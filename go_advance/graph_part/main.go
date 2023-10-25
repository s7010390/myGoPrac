package main

import "fmt"

type Graph map[string][]string

func DFS(graph Graph, node string, visited map[string]bool) {
	if visited[node] {
		return
	}

	fmt.Printf("%s -> ", node)
	visited[node] = true

	for _, neighbor := range graph[node] {
		DFS(graph, neighbor, visited)
	}
}

func main() {
	graph := make(Graph)
	graph["A"] = []string{"B", "C"}
	graph["B"] = []string{"A", "D", "E"}
	graph["C"] = []string{"A", "F"}
	graph["D"] = []string{"B"}
	graph["E"] = []string{"B", "F"}
	graph["F"] = []string{"C", "E"}

	startNode := "A"
	visited := make(map[string]bool)

	fmt.Printf("Depth-First Search starting from node %s:\n", startNode)
	DFS(graph, startNode, visited)
	fmt.Println()
}
