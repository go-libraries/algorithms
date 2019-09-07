/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package main

func InsertionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}

	}
}
func InsertionSortAdvance(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		e := arr[i]
		j := i
		for ; j > 0 && arr[j-1] > e; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}
