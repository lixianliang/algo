package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveK(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	i := 1
	one := head
	for {
		if one == nil {
			// len < k
			return nil
		}

		if i == k {
			break
		}

		i++
		one = one.Next
	}

	if one.Next == nil {
		return head.Next
	} else {
		prev := head
		one := one.Next
		two := head.Next

		for {
			if one.Next == nil {
				break
			}

			one = one.Next
			two = two.Next
			prev = prev.Next
		}

		prev.Next = two.Next
	}

	return head
}

func main() {
	head := ListNode{}
	RemoveK(&head, 1)
}
