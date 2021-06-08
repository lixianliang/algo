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
	if head == nil {
		return []int{}
	} else {
		return append(reversePrint(head.Next), head.Val)
	}
}

func main() {
	fmt.Println("", reversePrint(nil))
}
