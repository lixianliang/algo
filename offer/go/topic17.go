package main

import (
	"fmt"
)

func printNumbers(n int) []int {
	max := 0
	for i := 0; i < n; i++ {
		max = max*10 + 9
	}

	arr := make([]int, max, max)

	for i := 1; i <= max; i++ {
		arr[i-1] = i
	}

	return arr
}

func main() {
	fmt.Println(printNumbers(3))
	fmt.Println(printNumbers(999))
}
