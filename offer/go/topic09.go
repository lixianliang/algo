package main

import (
	"fmt"
)

type CQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() CQueue {
	return CQueue{stack1: make([]int, 0, 1000), stack2: make([]int, 0, 1000)}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	val := -1
	if len(this.stack2) == 0 {
		if len(this.stack1) == 0 {
		} else if len(this.stack1) == 1 {
			val = this.stack1[0]
			this.stack1 = []int{}
		} else {
			for i := len(this.stack1) - 1; i > 0; i-- {
				this.stack2 = append(this.stack2, this.stack1[i])
			}
			val = this.stack1[0]
			this.stack1 = []int{}
		}
	} else {
		val = this.stack2[len(this.stack2)-1]
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}
	return val
}

/**
* Your CQueue object will be instantiated and called as such:
* obj := Constructor();
* obj.AppendTail(value);
* param_2 := obj.DeleteHead();
 */

func main() {
	q := Constructor()
	q.AppendTail(3)
	fmt.Println(q.DeleteHead(), q.DeleteHead())
}
