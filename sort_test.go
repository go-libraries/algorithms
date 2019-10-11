package main

import "testing"

func TestSort(t *testing.T) {
	n := 10000
	s1 := generateRandomArray(n, 0, n)
	s2 := copyArray(s1, n)
	s3 := copyArray(s1, n)
	s4 := copyArray(s1, n)
	timeSpent("MergeSort", MergeSort, s3)
	timeSpent("QuickSort", QuickSort, s4)
	timeSpent("InsertionSortAdvance", InsertionSortAdvance, s1)
	timeSpent("SelectionSort", SelectionSort, s2)
}
