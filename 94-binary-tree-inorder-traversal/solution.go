package binarytreeinordertraversal

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	l := inorderTraversal(root.Left)
	r := inorderTraversal(root.Right)
	res := append(l, root.Val)
	res = append(res, r...)
	return res
}
