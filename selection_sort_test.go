package main

import "testing"

func TestSelectionSort(t *testing.T) {
	n := 100000
	slice := generateRandomArray(n, 0, n)
	timeSpent("SelectionSort", SelectionSort, slice)
}
