package main

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	n := 500000
	s1 := generateRandomArray(n, 0, n)
	//var s2 []int
	var s2 = make([]int, n)
	copy(s2, s1)

	timeSpent("InsertionSort", InsertionSort, s1)

	timeSpent("InsertionSortAdvance", InsertionSortAdvance, s2)
}
func TestInsertionSortAdvance(t *testing.T) {
	n := 100000
	slice := generateRandomArray(n, 0, n)
	timeSpent("InsertionSortAdvance", InsertionSortAdvance, slice)
}
