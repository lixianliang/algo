package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func FindFirstCommonNode(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return nil
	}

	// 将head1收尾相连，起来就是求链表是否有环的问题
	tmp := head1
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = head1

	// 判断是否有环
	hasLoop := false
	slow := head2
	fast := head2
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			hasLoop = true
			break
		}
	}

	// 求环的入口节点
	if hasLoop {
		slow2 := head2
		for slow != slow2 { // 开始就判断，有可能头节点就是环的入口
			slow = slow.Next
			slow2 = slow2.Next
		}
		tmp.Next = nil // 需要做对应的重置
		return slow
	} else {
		return nil
	}

	return nil
}

func main() {
	var x map[*ListNode]bool
	fmt.Println(x)
}
