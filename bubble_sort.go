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
func BubbleSortAdvance(arr []int) {
	n := len(arr)
	var newn int
	for {
		newn = 0
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				newn = i
			}
		}
		n = newn
		if newn <= 0 {
			break
		}
	}
}
