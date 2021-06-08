package main

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	// 没有考虑负整数的情况
	count := 0
	for num > 0 { // num > 0要改成num，如果是负数
		if num%2 == 1 {
			count++
		}
		// or count += num & 1
		num >>= 1
	}
	return count
}

func hammingWeight2(num int32) int {
	// 没有考虑负整数的情况 复现确实有问题，出现死循环
	count := 0
	for num != 0 { // num > 0要改成num，如果是负数
		if num%2 == 1 {
			count++
		}
		// or count += num & 1
		num >>= 1
	}
	return count
}

func hammingWeight3(num int32) int {
	count := 0
	flag := int32(1) // 通过来做判断和左移 需要理解负数右移会在最右边添加1
	for flag != 0 {  // num > 0要改成num，如果是负数
		if (num & flag) > 0 {
			count++
		}
		flag <<= 1
	}
	return count
}

func main() {
	//fmt.Println(hammingWeight2(-1245))
	fmt.Println(hammingWeight3(-1245))
}
