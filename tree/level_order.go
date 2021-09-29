package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0, 8)
	nodes := make([]*TreeNode, 0, 8)
	nodes = append(nodes, root)
	low, high := 0, len(nodes)

	for low != high {
		count := 0
		val := make([]int, 0, 8)
		for i := low; i < high; i++ {
			val = append(val, nodes[i].Val)
			if nodes[i].Left != nil {
				count++
				nodes = append(nodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				count++
				nodes = append(nodes, nodes[i].Right)
			}
		}
		low = high
		high = high + count
		res = append(res, val)
	}

	return res
}

func main() {

}
