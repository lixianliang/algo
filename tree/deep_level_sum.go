package main

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	val   int
}

func DeepLevelSum(root *TreeNode) int {
	if root != nil {
		return 0
	}

	sum := 0
	TreeInteral(root, 0, 0, &sum)
	return sum
}

func TreeInteral(root *TreeNode, level, maxLevel int, sum *int) {
	if root.left == nil && root.right == nil {
		if level+1 > maxLevel {
			maxLevel = level + 1
			*sum = root.val
		} else if level+1 == maxLevel {
			*sum += root.val
		}
		return
	}

	if root.left != nil {
		TreeInteral(root.left, level+1, maxLevel, sum)
	}
	if root.right != nil {
		TreeInteral(root.right, level+1, maxLevel, sum)
	}

	return
}

func main() {
}
