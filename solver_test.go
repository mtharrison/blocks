package main

import (
	"reflect"
	"testing"
)

func TestSortedMapKeys(t *testing.T) {
	m := map[int]bool{
		9: true,
		7: true,
		4: true,
	}

	keys := SortedMapKeys(m)
	expected := []int{4, 7, 9}

	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("Incorrect order for keys %v expected %v", keys, expected)
	}
}

func TestSolve(t *testing.T) {

	puzzle := Puzzle{
		Data: []int{
			0, 0, 1, 2,
			0, 1, 2, 1,
			2, 1, 1, 1,
		},
		Size: 4,
	}

	solution := Solve(puzzle)

	expected := []int{5, 9, 10, 11, 7}

	if !reflect.DeepEqual(solution, expected) {
		t.Errorf("Incorrect solution %v expected %v", solution, expected)
	}
}
