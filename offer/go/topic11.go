package main

import "fmt"

func minArray1(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		if i != 0 {
			if numbers[i] < numbers[i-1] {
				return numbers[i]
			}
		}
	}
	return numbers[0]
}

func minArray(numbers []int) int {
	var min, max = 0, len(numbers) - 1
	index := 0

	for numbers[min] >= numbers[max] {
		if min+1 == max {
			index = max
			break
		}

		mid := (min + max) / 2

		// 三个数都相等，需要顺序查找
		if numbers[min] == numbers[mid] && numbers[min] == numbers[max] {
			for i := 0; i < len(numbers); i++ {
				if i != 0 {
					if numbers[i] < numbers[i-1] {
						return numbers[i]
					}
				}
			}
			return numbers[0]
		}

		if numbers[mid] >= numbers[min] {
			// 前面子数组
			min = mid
		} else if numbers[mid] <= numbers[max] {
			// 后面子数组
			max = mid
		}
	}

	return numbers[index]
}

func main() {
	fmt.Println(minArray([]int{2, 2, 2, 0, 1}))
}
