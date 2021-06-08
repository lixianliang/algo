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

	res := 1
	for n > 4 {
		res = res * 3 % 1000000007
		n -= 3
	}

	return res * n % 1000000007
}

func main() {
	fmt.Println(cuttingRope(120))
}
