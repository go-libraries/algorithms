package main

import "testing"

func TestBubbleSort(t *testing.T) {
	n := 100000
	s1 := generateRandomArray(n, 0, n)
	var s2 = make([]int, n)
	copy(s2, s1)
	timeSpent("BubbleSort", BubbleSort, s1)
	timeSpent("BubbleSort", BubbleSortAdvance, s2)
}
