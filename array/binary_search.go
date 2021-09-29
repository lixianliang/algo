package main

import (
	"fmt"
)

func BinarySearch(arr []int, x int) int {
	if len(arr) < 1 {
		return -1
	}

	start := 0
	end := len(arr) - 1

	//for end > start {	// 	xuyao dengyu
	for start <= end {
		// mid := (start + end) / 2
		mid := start + (end-start)/2 // baozheng mid buhuiyichu
		fmt.Println(mid)
		if x < arr[mid] {
			end = mid - 1
		} else if x > arr[mid] {
			start = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	fmt.Print(BinarySearch([]int{1, 3, 5, 6, 8, 9, 10, 11, 15}, 9))
}
