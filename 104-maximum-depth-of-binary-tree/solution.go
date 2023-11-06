package maximumdepthofbinarytree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
