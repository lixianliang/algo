package main

import (
	"fmt"
)

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	QuickSortInternal(arr, 0, len(arr)-1)
}

func QuickSortInternal(arr []int, start, end int) {
	if start >= end {
		return
	}

	p := Partition(arr, start, end)
	QuickSortInternal(arr, start, p-1)
	QuickSortInternal(arr, p+1, end)
}

func Partition(arr []int, start, end int) int {
	pivot := arr[end]
	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot { // 小于
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}

func main() {
	arr := []int{6, 9, 4, 8, 1}
	QuickSort(arr)
	fmt.Println(arr)
}
