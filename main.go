package main

import "fmt"

type BlockNode struct {
	id         int
	color      int
	neighbours map[int]bool
}

func main() {
	input := []int{
		0, 0, 1, 2,
		0, 1, 2, 1,
		2, 1, 1, 1,
	}

	rowSize := 4

	fmt.Println(solve(input, rowSize))
}

func solve(input []int, rowSize int) []int {

	// Build the block graph without neighbour info

	graph := []BlockNode{}

	for i, v := range input {
		node := BlockNode{id: i, color: v, neighbours: make(map[int]bool)}
		graph = append(graph, node)
	}

	// Fill neighbour info

	for i, b := range graph {

		potentialNeighbours := []BlockNode{}

		if i%rowSize != 0 {
			potentialNeighbours = append(potentialNeighbours, graph[i-1])
		}

		if i%rowSize != rowSize-1 {
			potentialNeighbours = append(potentialNeighbours, graph[i+1])
		}

		if i/rowSize != 0 {
			potentialNeighbours = append(potentialNeighbours, graph[i-rowSize])
		}

		if i/rowSize != (len(graph)/rowSize)-1 {
			potentialNeighbours = append(potentialNeighbours, graph[i+rowSize])
		}

		for _, n := range potentialNeighbours {
			if n.color == b.color {
				b.neighbours[n.id] = true
			}
		}
	}

	// Find the largest component via depth first traversal

	largest := []int{}

	for i := range graph {
		visited := make(map[int]bool)
		dfs(i, graph, visited)
		if len(visited) > len(largest) {
			largest = []int{}
			for k := range visited {
				largest = append(largest, k)
			}
		}
	}

	return largest
}

func dfs(i int, g []BlockNode, visited map[int]bool) {

	// Check if already visited

	_, ok := visited[i]
	if ok {
		return
	}

	// Set current node visited

	node := g[i]
	visited[i] = true

	// Recurse on neighbours

	for i := range node.neighbours {
		dfs(i, g, visited)
	}
}
