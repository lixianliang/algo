package main

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	var rows, cols = len(board), len(board[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, i, j, k int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}

	// 省略了记录装填的数据结构
	board[i][j] = '0'
	// 上下左右
	res := dfs(board, word, i-1, j, k+1) || dfs(board, word, i+1, j, k+1) || dfs(board, word, i, j-1, k+1) || dfs(board, word, i, j+1, k+1)
	board[i][j] = word[k]
	return res
}

func main() {
	m := [][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")}
	w := "ABCCED"
	fmt.Println(exist(m, w))
}
