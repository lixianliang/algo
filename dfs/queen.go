package main

import (
	"fmt"
)

func solveNQueens(n int) [][]string {
	// 使用二维数组存储状态
	boards := make([][]int, n, n)
	for i := 0; i < n; i++ {
		boards[i] = make([]int, n, n)
	}

	Queue(boards, 0)

	return nil
}

func Queue(boards [][]int, index int) {
	n := len(boards)
	if index == len(boards) {
		count++
		fmt.Println("count: ", count)
		//Print(boards)
		//fmt.Println(boards)
		for _, arr := range boards {
			fmt.Println(arr)
		}
		return
	}

	for i := 0; i < n; i++ {
		if Check(boards, index, i) {
			//		fmt.Println(index, i)
			boards[index][i] = 1
			Queue(boards, index+1)
			boards[index][i] = 0
		}
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func Check(boards [][]int, row, col int) bool {
	// 列检查
	for i := 0; i < row; i++ {
		if boards[i][col] == 1 {
			return false
		}
	}

	// 斜对角检查
	for i := 0; i < row; i++ {
		//for j := 0; j < col; j++ {	// 这里列应该是n，不然会检查不全
		for j := 0; j < len(boards); j++ {
			x1 := row - i
			x2 := col - j
			if Abs(x1) == Abs(x2) && boards[i][j] == 1 { // 判断两者的绝对值是否相等
				return false
			}
		}
	}

	return true
}

func Print(boards [][]int) {
	n := len(boards)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if boards[i][j] == 1 {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println("")
	}
}

var count int = 0

func main() {
	solveNQueens(8)
}
