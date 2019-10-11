/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package main

func ShellSort(arr []int) {
	n := len(arr)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			e := arr[i]
			j := i
			for ; j >= h && e < arr[j-h]; j -= h {
				arr[j] = arr[j-h]
			}
			arr[j] = e
		}
		h /= 3
	}
}
