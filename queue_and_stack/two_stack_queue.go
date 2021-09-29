package main

var stack1 []int
var stack2 []int

func Push(val int) {
	if len(stack2) != 0 {
		for i := len(stack2) - 1; i >= 0; i-- {
			stack1 = append(stack1, stack2[i])
		}
		stack2 = []int{}
	}

	stack1 = append(stack1, val)
}

func Pop() int {
	if len(stack1) == 0 && len(stack2) == 0 {
		return -1
	}

	if len(stack1) > 0 {
		for i := len(stack1) - 1; i >= 0; i-- {
			stack2 = append(stack2, stack1[i])
		}
		stack1 = []int{}
	}
	x := stack2[len(stack2)-1]
	stack2 = stack2[:len(stack2)-1]

	return x
}

func main() {
}
