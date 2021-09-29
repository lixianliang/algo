package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0, 4)
	path := make([]int, 0, 4)

	pathSumInternal(root, path, 0, sum, &res)
	return res
}

func pathSumInternal(root *TreeNode, path []int, cur, sum int, res *[][]int) {
	cur += root.Val
	/*if cur > sum {
		return
	}*/

	path = append(path, root.Val)

	if root.Left == nil && root.Right == nil {
		if cur == sum {
			tmp := make([]int, len(path)) // 这里需要做一次copy，避免互相影响
			copy(tmp, path)
			*res = append(*res, tmp)
		}
	} else {
		if root.Left != nil {
			pathSumInternal(root.Left, path, cur, sum, res)
		}
		if root.Right != nil {
			pathSumInternal(root.Right, path, cur, sum, res)
		}
	}
	path = path[:len(path)-1] // 这一行需要加上，如果回溯到上一层，则将当前的path值去掉
}

func main() {
}
