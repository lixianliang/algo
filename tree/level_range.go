package main

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	brt   *TreeNode
}

func TreeRange(tree *TreeNode) {
	// 自己的左节点指向右边节点
	if tree.left != nil && tree.right != nil {
		tree.left.brt = tree.right
	}
	// 右边节点指向兄弟的左节点
	if tree.right != nil && tree.brt != nil && tree.brt.left != nil {
		tree.right.brt = tree.brt.left
	}

	if tree.left != nil {
		TreeRange(tree.left)
	}
	if tree.right != nil {
		TreeRange(tree.right)
	}
}

func main() {
	var tree TreeNode
	TreeRange(&tree)
}
