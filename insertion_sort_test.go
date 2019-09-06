package main

import (
	"testing"
)

func TestSort(t *testing.T) {
	n := 100000
	slice := generateRandomArray(n, 0, n)
	timeSpent("InsertionSort", insertion_sort, slice)
	//slice2 := generateRandomArray(n, 0, n)
	//timeSpent("InsertionSort", InsertionSort.Sort, slice2)
}
