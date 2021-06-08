package main

import (
	"fmt"
)

func findRepeatNumber(nums []int) int {
	numHash := make(map[int]int, len(nums))
	for _, n := range nums {
		if _, ok := numHash[n]; ok {
			return n
		} else {
			numMap[n] = 1
		}
	}
	return -1
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println("find", findRepeatNumber(nums))
}
