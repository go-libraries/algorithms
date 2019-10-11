package main

import "testing"

func TestBubbleSort(t *testing.T) {
	n := 10000
	s1 := generateRandomArray(n, 0, n)
	s2 := copyArray(s1, n)

	timeSpent("BubbleSort", BubbleSort, s1)
	timeSpent("BubbleSort", BubbleSortAdvance, s2)
}
