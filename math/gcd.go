package main

import (
	"fmt"
)

func gcd(a, b int) int {
	if a < b { // 这里去掉也没关系 后续计算会旋正过来
		a, b = b, a
	}
	c := a % b

	for c != 0 {
		a = b
		b = c
		c = a % b
	}

	return b
}

func main() {
	fmt.Println(gcd(8, 12))
	//	fmt.Println(8 % 12)
}
