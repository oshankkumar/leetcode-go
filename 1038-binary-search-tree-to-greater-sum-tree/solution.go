package binarysearchtreetogreatersumtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var _ = bstToGst

func _bstToGST(root *TreeNode, runningSum int) (*TreeNode, int) {
	if root == nil {
		return nil, runningSum
	}

	root.Right, runningSum = _bstToGST(root.Right, runningSum)
	runningSum += root.Val
	root.Val = runningSum
	root.Left, runningSum = _bstToGST(root.Left, runningSum)
	return root, runningSum
}

func bstToGst(root *TreeNode) *TreeNode {
	root, _ = _bstToGST(root, 0)
	return root
}
