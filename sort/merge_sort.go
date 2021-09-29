package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortInList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tmp := slow.Next
	slow.Next = nil

	head1 := sortInList(head)
	head2 := sortInList(tmp)

	return Merge(head1, head2)
}

func Merge(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	if head1.Val < head2.Val {
		head1.Next = Merge(head1.Next, head2)
		return head1
	} else {
		head2.Next = Merge(head1, head2.Next)
		return head2
	}
}

func main() {

}
