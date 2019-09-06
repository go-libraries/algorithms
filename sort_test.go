package main

import (
	"Algorithms/InsertionSort"
	"Algorithms/SelectionSort"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func generateRandomArray(n, min, max int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		num := rand.Intn(max-min) + min
		arr = append(arr, num)
	}
	return arr
}

// 判断arr数组是否有序
func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
func timeSpent(funcName string, inner func(arr []int), arr []int) {
	start := time.Now()
	inner(arr)
	fmt.Println(funcName+" time spent:", time.Since(start).Seconds())

}
func TestSort(t *testing.T) {
	n := 100000
	slice := generateRandomArray(n, 0, n)
	timeSpent("SelectionSort", SelectionSort.Sort, slice)
	slice2 := generateRandomArray(n, 0, n)
	timeSpent("InsertionSort", InsertionSort.Sort, slice2)
}
