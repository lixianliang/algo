package main

import (
	"fmt"
)

func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	/* 还可以将数组省略，通过val1 val2 val3
	val3 = val1 + val2
	val1 = val2 % 1000000007
	val2 = val3 % 1000000007
	*/
	var val1, val2, val3 int
	val1 = 1
	val2 = 2
	for i := 3; i <= n; i++ {
		val3 = val1 + val2
		val1 = val2 % 1000000007
		val2 = val3 % 1000000007
	}
	return val2
}

func main() {
	fmt.Println(numWays(7))
}
