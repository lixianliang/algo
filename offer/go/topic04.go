package main

import (
	"fmt"
)

func findNumberIn2DArray(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	col := len(matrix[0])
	if col == 0 {
		return false
	}
	var curRow, curClo = 0, col - 1
	for curRow < row && curClo >= 0 {
		if target > matrix[curRow][curClo] {
			curRow++
		} else if target < matrix[curRow][curClo] {
			curClo--
		} else {
			return true
		}
	}

	return false
}

func main() {
	var matrix = [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(findNumberIn2DArray(matrix, 5))
	fmt.Println(findNumberIn2DArray(matrix, 20))
}
