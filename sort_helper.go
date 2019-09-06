package main

import (
	"fmt"
	"math/rand"
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
