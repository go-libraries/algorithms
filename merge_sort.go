/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package main

func MergeSort(arr []int) {
	n := len(arr)
	sort(arr, 0, n-1)
}
func sort(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	sort(arr, left, mid)
	sort(arr, mid+1, right)
	merge(arr, left, mid, right)
}
func merge(arr []int, left, mid, right int) {

}
