package minimumdepthofbinarytree

const Max = 1<<31 - 1

func minDepth(root *TreeNode) int {
	if d := minimumDepth(root, 0, Max); d < Max {
		return d
	}
	return 0
}

func minimumDepth(root *TreeNode, prevHeight int, minDepthTillNow int) int {
	if root == nil {
		return minDepthTillNow
	}

	nodeHeight := prevHeight + 1
	if isLeaf(root) && nodeHeight < minDepthTillNow {
		minDepthTillNow = nodeHeight
	}

	minDepthTillNow = minimumDepth(root.Left, nodeHeight, minDepthTillNow)
	minDepthTillNow = minimumDepth(root.Right, nodeHeight, minDepthTillNow)

	return minDepthTillNow
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
