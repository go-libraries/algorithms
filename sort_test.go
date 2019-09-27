package main

import "testing"

func TestSort(t *testing.T) {
	n := 1000000
	s1 := generateRandomArray(n, 0, n)
	var s2 = make([]int, n)
	var s3 = make([]int, n)
	var s4 = make([]int, n)
	copy(s2, s1)
	copy(s3, s1)
	copy(s4, s1)
	timeSpent("MergeSort", MergeSort, s3)
	timeSpent("QuickSort", QuickSort, s4)
	timeSpent("InsertionSortAdvance", InsertionSortAdvance, s1)
	timeSpent("SelectionSort", SelectionSort, s2)
}
