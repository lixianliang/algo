package main

import (
	"fmt"
)

// 超出时间限制
/*func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return (fib(n-1) + fib(n-2)) % 1000000007
	}
}*/

func fib(n int) int {
	var arr [101]int
	arr[0] = 0
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = (arr[i-1] + arr[i-2]) % 1000000007
	}
	return arr[n]
	/* 还可以将数组省略，通过val1 val2 val3
	val3 = val1 + val2
	va1 = val2 % 1000000007
	val2 = val3 % 1000000007
	*/
}

func main() {
	fmt.Println(fib(95))
}
