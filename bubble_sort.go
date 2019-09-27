package main

func BubbleSort(arr []int) {
	n := len(arr)
	swapped := false
	for {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
		n--
		if swapped == false {
			break
		}
	}

}
