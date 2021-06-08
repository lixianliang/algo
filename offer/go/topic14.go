package main

import (
	"fmt"
)

func cuttingRope(n int) int {
	if n < 2 {
		return 0
	} else if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}

	arr := make([]int, n+1, n+1)
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3

	for i := 4; i <= n; i++ {
		max := 0
		for j := 1; j <= i/2; j++ {
			x := arr[j] * arr[i-j]
			if x > max {
				max = x
			}
		}
		arr[i] = max
	}

	return arr[n]
}

func main() {
	fmt.Println(cuttingRope(10))
}
