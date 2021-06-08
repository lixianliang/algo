package main

import (
	"fmt"
)

func getSum(x int) int {
	res := 0
	for ; x > 0; x /= 10 {
		res += x % 10
	}
	return res
}

func movingCount(m int, n int, k int) int {
	mat := make([][]int, m, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]int, n, n)
	}
	/*for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = 0
		}
	}*/

	count := 0
	dfs(0, 0, k, mat, &count)

	return count
}

func dfs(i, j, k int, m [][]int, count *int) bool {
	if i < 0 || i >= len(m) || j < 0 || j >= len(m[0]) || m[i][j] == 1 {
		return false
	}

	m[i][j] = 1
	if getSum(i)+getSum(j) <= k {
		*count += 1
		dfs(i-1, j, k, m, count)
		dfs(i+1, j, k, m, count)
		dfs(i, j-1, k, m, count)
		dfs(i, j+1, k, m, count)
		return true
	} else {
		return false
	}
}

func main() {
	//fmt.Println(movingCount(2, 3, 1))
	fmt.Println(movingCount(7, 2, 3))
}
