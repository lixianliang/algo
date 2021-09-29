package main

import (
	"fmt"
)

var sum int = 0

func ListAll(arr []int) {
	if len(arr) == 0 {
		return
	}

	partion(arr, 0)
}

func partion(arr []int, start int) {
	if start == len(arr) {
		sum++
		fmt.Println(arr)

		return
	}

	var prev int
	prev = arr[start]
	for i := start; i < len(arr); i++ {
		if i != start {
			// 判断是否和前驱数字一样
			if prev == arr[i] {
				break
			}
		}
		arr[i], arr[start] = arr[start], arr[i]
		partion(arr, start+1)
		arr[i], arr[start] = arr[start], arr[i]
		prev = arr[i]
	}
}

func isswap(arr []int, index int) bool {
	for i := index + 1; i < len(arr); i++ {
		if arr[index] == arr[i] {
			return false
		}
	}
	// 判断与后续数字没有重复才进行交换
	return true
}

func dfs(arr []int, start int, arr2 [][]int) [][]int {
	if start == len(arr) {
		fmt.Println(arr)
		arr2 = append(arr2, arr)
		return arr2
	}

	for i := start; i < len(arr); i++ {
		if isswap(arr, i) { // 判断是否交换处理
			arr[i], arr[start] = arr[start], arr[i]
			arr2 = dfs(arr, start+1, arr2)
			arr[i], arr[start] = arr[start], arr[i]
		}
	}

	return arr2
}

func permuteUnique(num []int) [][]int {
	// write code here
	arr2 := [][]int{}
	return dfs(num, 0, arr2)
}

func main() {
	//ListAll([]int{1, 2, 3, 3})
	//ListAll([]int{1, 2, 2})
	//ListAll([]int{2, 1, 2}) // 需要排序
	//fmt.Println(sum)
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
