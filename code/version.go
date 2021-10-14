package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EqualVersion(v1, v2 string) int {
	arr1 := strings.Split(v1, ".")
	arr2 := strings.Split(v2, ".")

	l1, l2 := len(arr1), len(arr2)
	index := l1
	if l1 > l2 {
		index = l2
	}
	for i := 0; i < index; i++ {
		x1, _ := strconv.ParseUint(arr1[i], 10, 32)
		x2, _ := strconv.ParseUint(arr2[i], 10, 32)
		if x1 > x2 {
			return 1
		} else if x1 < x2 {
			return -1
		}
	}

	if l1 > l2 {
		for i := l2; i < l1; i++ {
			x, _ := strconv.ParseUint(arr1[i], 10, 32)
			if x > 0 {
				return 1
			}
		}
	} else if l1 < l2 {
		for i := l1; i < l2; i++ {
			x, _ := strconv.ParseUint(arr2[i], 10, 32)
			if x > 0 {
				return -1
			}
		}
	}

	return 0
}

func main() {
	fmt.Println(EqualVersion("2.2.1", "2.2"))
	fmt.Println(EqualVersion("3.5.2", "4.0"))
}
