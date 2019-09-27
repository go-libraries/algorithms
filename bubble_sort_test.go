package main

import "testing"

func TestBubbleSort(t *testing.T) {
	n := 1000000
	s1 := generateRandomArray(n, 0, n)
	timeSpent("BubbleSort", BubbleSort, s1)

}
