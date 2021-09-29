package main

import (
	"fmt"
)

func GetUglyNumber(index int) int {
	if index < 1 {
		//return -1
		return 0 // 要返回为0
	}

	dp := make([]int, index, index)
	dp[0] = 1 // 1
	i2, i3, i5, i := 0, 0, 0, 1
	for i < index {
		dp[i] = min(min(dp[i2]*2, dp[i3]*3), dp[i5]*5)
		// 逻辑错误，里面可能有重复的，需要进行过滤
		/*if dp[i] == dp[i2]*2 {
			i2++
		} else if dp[i] == dp[i3]*3 {
			i3++
		} else if dp[i] == dp[i5]*5 {
			i5++
		}
		fmt.Println(dp[i], i2, i3, i5)*/
		if dp[i] == dp[i2]*2 {
			i2++
		}
		if dp[i] == dp[i3]*3 {
			i3++
		}
		if dp[i] == dp[i5]*5 {
			i5++
		}
		i++
	}

	//fmt.Println(dp)
	return dp[index-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(GetUglyNumber(10))
}
