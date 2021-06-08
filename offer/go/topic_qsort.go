package main

import (
	"fmt"
)

// jike时间使用的是单向扫描，个人感觉双向扫描更好理解
// 以维度元素为基准
// 如果i j相等可以不进行swap
/*
partition(arr, p, r) {
	pivot := arr[r]
	i := p
	for j := p to r-1 do {
		if arr[j] < pivot {
			swap arr[i] whit arr[j]
			i := i + 1
		}
	}
	swap arr[i] with arr[r]
	return i
}
*/

// 双向扫描，以初始元素为基准
// 设置i j先从j从右往左找小于pivot i从做忘右大于pivot 然后进行swap

func Partition(arr []int, start, end int) int {
	pivot := arr[end]
	i := start
	for j := start; j < end-1; j++ {
		if arr[j] < pivot {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]

	return i
}

func QuickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	q := Partition(arr, start, end)
	QuickSort(arr, start, q-1)
	QuickSort(arr, q+1, end)
}

func main() {
	arr := []int{3, 2, 5, 9, 4}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
