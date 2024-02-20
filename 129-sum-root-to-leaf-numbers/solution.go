package sumroottoleafnumbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return sumNumbersUtil(root, 0)
}

func sumNumbersUtil(root *TreeNode, mul int) int {
	if root == nil {
		return 0
	}

	placeVal := mul*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return placeVal
	}

	return sumNumbersUtil(root.Left, placeVal) + sumNumbersUtil(root.Right, placeVal)
}
