/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package main

func MergeSort(arr []int) {
	mergesort(arr, 0, len(arr)-1)
}
func mergesort(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	mergesort(arr, left, mid)
	mergesort(arr, mid+1, right)
	merge(arr, left, mid, right)
}
func merge(arr []int, left, mid, right int) {
	var tmp = make([]int, right-left+2)
	copy(tmp, arr[left:right+1])

	i := left
	j := mid + 1

	for k := left; k <= right; k++ {
		if i > mid {
			arr[k] = tmp[j-left]
			j++
		} else if j > right {
			arr[k] = tmp[i-left]
			i++
		} else if tmp[i-left] < tmp[j-left] {
			arr[k] = tmp[i-left]
			i++
		} else {
			arr[k] = tmp[j-left]
			j++
		}
	}
}
