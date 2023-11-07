package maximumbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var _ = constructMaximumBinaryTree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return constructMaximumBinaryTreeUtil(nums, 0, len(nums)-1)
}

func constructMaximumBinaryTreeUtil(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	maxIdx := maxIndex(nums, start, end)

	root := &TreeNode{Val: nums[maxIdx]}
	root.Left = constructMaximumBinaryTreeUtil(nums, start, maxIdx-1)
	root.Right = constructMaximumBinaryTreeUtil(nums, maxIdx+1, end)

	return root
}

func maxIndex(nums []int, start, end int) int {
	maxPos := start
	for i := start + 1; i <= end; i++ {
		if nums[i] > nums[maxPos] {
			maxPos = i
		}
	}
	return maxPos
}
