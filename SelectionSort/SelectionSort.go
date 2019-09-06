package main

import "fmt"

func sort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		tmp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = tmp
	}
}
func main() {
	arr := [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort(arr[:])
	fmt.Println(arr)
}
