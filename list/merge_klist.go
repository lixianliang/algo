package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	// 需要做非空判断
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	// i+1要小于len
	var i int
	for i = 0; i+1 < len(lists); i += 2 {
		x1 := Merge(lists[i], lists[i+1])
		lists = append(lists, x1)
	}

	// 返回最后一个list
	return lists[len(lists)-1]
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
