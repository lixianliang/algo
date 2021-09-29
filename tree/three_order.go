package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func threeOrders(root *TreeNode) [][]int {
	if root == nil {
		// return nil
		return [][]int{{}, {}, {}} // 需要题目要求
	}

	res := make([][]int, 0, 3)

	arr1 := make([]int, 0, 8)
	PreOrder(root, &arr1)
	res = append(res, arr1)

	arr2 := make([]int, 0, 8)
	MidOrder(root, &arr2)
	res = append(res, arr2)

	arr3 := make([]int, 0, 8)
	PostOrder(root, &arr3)
	res = append(res, arr3)

	return res
}

func PreOrder(root *TreeNode, arr *[]int) {
	*arr = append(*arr, root.Val)
	if root.Left != nil {
		PreOrder(root.Left, arr)
	}
	if root.Right != nil {
		PreOrder(root.Right, arr)
	}
}

func MidOrder(root *TreeNode, arr *[]int) {
	if root.Left != nil {
		MidOrder(root.Left, arr)
	}
	*arr = append(*arr, root.Val)
	if root.Right != nil {
		MidOrder(root.Right, arr)
	}
}

func PostOrder(root *TreeNode, arr *[]int) {
	if root.Left != nil {
		PostOrder(root.Left, arr)
	}
	if root.Right != nil {
		PostOrder(root.Right, arr)
	}
	*arr = append(*arr, root.Val)
}

func main() {
}
