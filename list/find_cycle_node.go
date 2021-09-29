package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func EntryNodeOfLoop(head *ListNode) *ListNode {

	hasLoop := false
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			hasLoop = true
			break
		}
	}

	if hasLoop {
		slow2 := head
		for slow != slow2 { // 开始就判断，有可能头节点就是环的入口
			slow = slow.Next
			slow2 = slow2.Next
		}
		return slow
	} else {
		return nil
	}
}

func main() {
}
