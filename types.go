package main

type Puzzle struct {
	Data []int
	Size int
}

type BlockNode struct {
	id         int
	color      int
	neighbours map[int]bool
}
