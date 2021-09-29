package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level := 0
	maxlevel := 0
	/*if root.Left != nil {	// 这里的逻辑可以放到递归里面
		treeDepth(root.Left, level+1, &maxlevel)
	}
	if root.Right != nil {
		treeDepth(root.Right, level+1, &maxlevel)
	}*/
	treeDepth(root, level+1, &maxlevel)

	return maxlevel
}

func treeDepth(root *TreeNode, level int, maxlevel *int) {
	if level > *maxlevel {
		*maxlevel = level
	}

	if root.Left != nil {
		treeDepth(root.Left, level+1, maxlevel)
	}
	if root.Right != nil {
		treeDepth(root.Right, level+1, maxlevel)
	}
}

func main() {
}
