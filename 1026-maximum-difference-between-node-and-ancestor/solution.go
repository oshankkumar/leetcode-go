package maximum_difference_between_node_and_ancestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsLeaf(t *TreeNode) bool {
	return t.Left == nil && t.Right == nil
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func maxAncestorDiff(root *TreeNode) int {
	return treeMaxDiff(root, (1<<31)-1, -1*(1<<31), 0)
}

func treeMaxDiff(root *TreeNode, currMin, currMax int, currMaxDiff int) int {
	if root == nil {
		return currMaxDiff
	}

	currMin = min(currMin, root.Val)
	currMax = max(currMax, root.Val)

	if IsLeaf(root) {
		currMaxDiff = max(currMaxDiff, abs(currMax, currMin))
	}

	currMaxDiff = treeMaxDiff(root.Left, currMin, currMax, currMaxDiff)
	currMaxDiff = treeMaxDiff(root.Right, currMin, currMax, currMaxDiff)

	return currMaxDiff
}
