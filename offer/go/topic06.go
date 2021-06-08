package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	arr := make([]int, 0, 10000)
	nodeArr := make([]*ListNode, 0, 10000)
	tmp := head
	for tmp != nil {
		nodeArr = append(nodeArr, tmp)
		tmp = tmp.Next
	}

	for i := len(nodeArr) - 1; i >= 0; i-- {
		arr = append(arr, nodeArr[i].Val)
	}

	return arr
}

func main() {
	fmt.Println("", reversePrint(nil))
}
