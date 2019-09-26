/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package main

func QuickSort(arr []int) {
	quicksort(arr, 0, len(arr)-1)
}
func quicksort(arr []int, left, right int) {
	if left >= right {
		return
	}
	p := partition(arr, left, right)
	quicksort(arr, left, p-1)
	quicksort(arr, p+1, right)
}
func partition(arr []int, left, right int) int {
	v := arr[left]
	j := left
	for i := left + 1; i <= right; i++ {
		if arr[i] < v {
			j++
			arr[i], arr[j] = arr[j], arr[i]
		}
		arr[left], arr[j] = arr[j], arr[left]
	}
	return j
}
