package secondminimumnodeinabinarytree

const Max = 1 << 31

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	res := findSecondMinimumValueUtil(root, root.Val, Max)
	if res < Max {
		return res
	}
	return -1
}

func findSecondMinimumValueUtil(root *TreeNode, min, secMin int) int {
	if root == nil {
		return secMin
	}

	if root.Val > min && root.Val < secMin {
		return root.Val
	}

	if root.Val == min {
		secMin = findSecondMinimumValueUtil(root.Left, min, secMin)
		secMin = findSecondMinimumValueUtil(root.Right, min, secMin)
	}

	return secMin
}
