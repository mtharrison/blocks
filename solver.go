package main

import (
	"sort"
)

type blockNode struct {
	id         int
	color      int
	neighbours map[int]bool
}

func solve(input []int, rowSize int) []int {

	// Build the block graph without neighbour info

	graph := []blockNode{}

	for i, v := range input {
		node := blockNode{id: i, color: v, neighbours: make(map[int]bool)}
		graph = append(graph, node)
	}

	// Fill neighbour info

	for i, b := range graph {

		neighbours := []int{}

		if i%rowSize != 0 {
			neighbours = append(neighbours, i-1)
		}

		if i%rowSize != rowSize-1 {
			neighbours = append(neighbours, i+1)
		}

		if i/rowSize != 0 {
			neighbours = append(neighbours, i-rowSize)
		}

		if i/rowSize != (len(graph)/rowSize)-1 {
			neighbours = append(neighbours, i+rowSize)
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
		nodes := dfs(i, graph, visited)
		if len(nodes) > len(largest) {
			largest = nodes
		}
	}

	return largest
}

func sortedMapKeys(m map[int]bool) []int {

	keys := []int{}
	for i := range m {
		keys = append(keys, i)
	}

	sort.Ints(keys)

	return keys
}

func dfs(i int, g []blockNode, visited map[int]bool) []int {

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

	for _, i := range sortedMapKeys(node.neighbours) {
		nodes = append(nodes, dfs(i, g, visited)...)
	}

	return nodes
}
