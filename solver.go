package main

import (
	"sort"
)

func Solve(p Puzzle) []int {

	// Build the block graph without neighbour info

	graph := []BlockNode{}

	for i, v := range p.Data {
		node := BlockNode{id: i, color: v, neighbours: make(map[int]bool)}
		graph = append(graph, node)
	}

	// Fill neighbour info

	for i, b := range graph {

		neighbours := []int{}

		if i%p.Size != 0 {
			neighbours = append(neighbours, i-1)
		}

		if i%p.Size != p.Size-1 {
			neighbours = append(neighbours, i+1)
		}

		if i/p.Size != 0 {
			neighbours = append(neighbours, i-p.Size)
		}

		if i/p.Size != (len(graph)/p.Size)-1 {
			neighbours = append(neighbours, i+p.Size)
		}

		for _, n := range neighbours {
			if graph[n].color == b.color {
				b.neighbours[n] = true
			}
		}
	}

	// Find the largest component via depth first traversal

	largest := []int{}
	visited := make(map[int]bool)

	for i := range graph {
		nodes := DepthFirstTraversal(i, graph, visited)
		if len(nodes) > len(largest) {
			largest = nodes
		}
	}

	return largest
}

func SortedMapKeys(m map[int]bool) []int {

	keys := []int{}
	for i := range m {
		keys = append(keys, i)
	}

	sort.Ints(keys)

	return keys
}

func DepthFirstTraversal(i int, g []BlockNode, visited map[int]bool) []int {

	// Check if already visited

	_, ok := visited[i]
	if ok {
		return []int{}
	}

	// Set current node visited

	node := g[i]
	visited[i] = true
	nodes := []int{i}

	// Recurse on neighbours

	for _, i := range SortedMapKeys(node.neighbours) {
		nodes = append(nodes, DepthFirstTraversal(i, g, visited)...)
	}

	return nodes
}
