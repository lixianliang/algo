package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func Merge(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	//var node *ListNode 优化node
	if head1.Val < head2.Val {
		//node = head1
		head1.Next = Merge(head1.Next, head2)
		return head1
	} else {
		//node = head2
		head2.Next = Merge(head1, head2.Next)
		return head2
	}
	//return node
}

func main() {
}
