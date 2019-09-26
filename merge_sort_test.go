/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package main

import "testing"

func TestMergeSort(t *testing.T) {
	n := 1000000
	slice := generateRandomArray(n, 0, n)
	timeSpent("MergeSort", MergeSort, slice)
}
