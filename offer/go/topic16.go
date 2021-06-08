package main

import (
	"fmt"
)

// 运算效率太低，需要改进为位运算
func myPow1(x float64, n int) float64 {
	// 注意负数
	if x == 0.0 && n < 0 {
		return 0.0
	}

	abs := n
	if n < 0 {
		abs = -n
	}

	count := 1.0
	for i := 0; i < abs; i++ {
		count *= x
	}

	if n < 0 {
		count = 1 / count
	}

	return count
}

func myPow(x float64, n int) float64 {
	if x == 0.0 && n < 0 {
		return 0.0
	}

	abs := n
	if n < 0 {
		abs = -n
	}

	res := myPowX(x, uint(abs))
	if n < 0 {
		res = 1 / res
	}

	return res
}

func myPowX(x float64, n uint) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	res := myPowX(x, n>>1)
	res *= res
	if (n & 1) == 1 {
		res *= x
	}

	return res
}

func main() {
	fmt.Println(myPow1(10, 3))
	fmt.Println(myPow1(0, -3))

	fmt.Println(myPow(10, 3))
	fmt.Println(myPow(2.0000, -2))
}
