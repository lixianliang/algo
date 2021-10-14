package main

import (
	"fmt"
)

func Count(arr []int) int {
	v := make([]int, len(arr))
	max := 0
	CountInternal(arr, v, 0, &max)
	return max
}

func CountInternal(arr []int, vist []int, index int, max *int) {
	if index == len(arr) {
		count := 0
		for i := 0; i < len(arr); i++ {
			if vist[i] == 1 {
				count += arr[i]
			}
		}
		if count > *max {
			*max = count
		}
	}

	for i := index; i < len(arr); i++ {
		vist[i] = 1
		if i != 0 {
			if vist[i-1] == 1 {
				vist[i] = 0
				continue
			}
		}
		CountInternal(arr, vist, index+1, max)
		vist[i] = 0
	}
}

func main() {
	//a := 0
	//fmt.Scan(&a)
	//fmt.Printf("%d\n", a)
	fmt.Printf("Hello World!\n")
	arr := []int{2, 7, 9, 3, 1}
	fmt.Println(Count(arr))
}
