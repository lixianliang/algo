package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func FindMiddleNode(head *ListNode) *ListNode {
	/*if head == nil {
		return nil
	}*/

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {
	a := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}
	b := &ListNode{Val: 3, Next: a}
	head := &ListNode{Val: 2, Next: b}
	head = nil
	mid := FindMiddleNode(head)
	fmt.Println(mid.Val)
}
