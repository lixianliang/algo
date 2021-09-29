package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := ReverseList(head.Next) // 得到已反转好的list

	head.Next.Next = head // 反转指向
	head.Next = nil

	return p
}

func main() {
}
