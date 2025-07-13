package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SearchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val < root.Val {
		return SearchBST(root.Left, val)
	} else if val > root.Val {
		return SearchBST(root.Right, val)
	}

	return root
}
