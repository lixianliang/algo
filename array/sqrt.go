package main

import (
	"fmt"
)

func sqrt(x int) int {
	low, high := 0, x // 平方根，结果肯定是正整数
	ans := -1

	for low <= high {
		mid := (low + high) / 2
		if mid*mid <= x { // 小于等于同一一个条件  向下取整
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans
}

func main() {
	fmt.Println(sqrt(3))
}
