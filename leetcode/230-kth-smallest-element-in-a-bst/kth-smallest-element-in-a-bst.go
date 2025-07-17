package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func KthSmallest(root *TreeNode, k int) int {
	resultOrder := []int{}

	var dfsRead func(root *TreeNode)
	dfsRead = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfsRead(root.Left)
		resultOrder = append(resultOrder, root.Val)
		dfsRead(root.Right)
	}

	dfsRead(root)

	return resultOrder[k-1]
}
