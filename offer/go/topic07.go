package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	node := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
			//rootCur := 0
			// 左子树
			/*if i-1 <= 0 {
				node.Left = nil
			} else {
				rootCur++
				node.Left = buildTree(preorder[rootCur:], inorder[0:i])
			}

			// 右子树
			if i+1 >= len(inorder)-1 {
				node.Right = nil
			} else {
				rootCur++
				node.Right = buildTree(preorder[rootCur:], inorder[i+1:])
			}*/
		}
	}
	node.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	node.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return node
}

func main() {
	buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
}
