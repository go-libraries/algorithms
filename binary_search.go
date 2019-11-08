package main

func binary_search(arr []int, n, target int) int {
	left := 0
	right := n
	for left < right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}
