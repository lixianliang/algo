package main

import (
	"fmt"
)

var sum int = 0

//var arr2 [][]int = [][]int{}

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

	for i := start; i < len(arr); i++ {
		arr[i], arr[start] = arr[start], arr[i]
		partion(arr, start+1)
		arr[i], arr[start] = arr[start], arr[i]
	}
}

func partion2(arr []int, start int, arr2 *[][]int) {
	if start == len(arr) {
		sum++
		//fmt.Println(arr)
		x := make([]int, len(arr), len(arr)) // 需要做对应数据的copy，不然有问题
		copy(x, arr)
		*arr2 = append(*arr2, x)
		return
	}

	for i := start; i < len(arr); i++ {
		arr[i], arr[start] = arr[start], arr[i]
		partion2(arr, start+1, arr2)
		arr[i], arr[start] = arr[start], arr[i]
	}
}

func permute(num []int) [][]int {
	// write code here
	//arr2 := [][]int{}
	arr2 := make([][]int, 0)
	partion2(num, 0, &arr2) // 传递参数的问题也可以使用匿名函数的方式来解决
	return arr2
}

func dfs(arr []int, start int, arr2 [][]int) [][]int {
	if start == len(arr) {
		sum++
		arr2 = append(arr2, arr)
		return arr2
	}

	for i := start; i < len(arr); i++ {
		arr[i], arr[start] = arr[start], arr[i]
		arr2 = dfs(arr, start+1, arr2)
		arr[i], arr[start] = arr[start], arr[i]
	}
	return arr2
}

func permute3(num []int) [][]int {
	// write code here
	//arr2 := [][]int{} var arr2 [][]int arr2 := make([][]int, 0) 都可以
	arr2 := [][]int{}
	arr2 = partion2(num, 0, arr2)
	return arr2
}

func main() {
	//ListAll([]int{1, 2, 3, 4})
	//fmt.Println(sum)
	fmt.Println(permute([]int{1, 2, 3, 4}))
	fmt.Println(sum)
}
