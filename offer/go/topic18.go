package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}

	prev := head
	tmp := head.Next
	for tmp != nil {
		if tmp.Val == val {
			prev.Next = tmp.Next
			break
		}
		prev = tmp
		tmp = tmp.Next
	}

	return head
}

func main() {

}
