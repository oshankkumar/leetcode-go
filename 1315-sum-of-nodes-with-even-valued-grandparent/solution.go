package sumofnodeswithevenvaluedgrandparent

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var _ = sumEvenGrandparent

func sumEvenGrandparentUtil(root *TreeNode, parents map[*TreeNode]*TreeNode, sum int) int {
	if root == nil {
		return sum
	}

	parent := parents[root]
	gp := parents[parent]

	if gp != nil && gp.Val&1 == 0 {
		sum += root.Val
	}

	parents[root.Left], parents[root.Right] = root, root
	sum = sumEvenGrandparentUtil(root.Left, parents, sum)
	sum = sumEvenGrandparentUtil(root.Right, parents, sum)
	return sum
}

func sumEvenGrandparent(root *TreeNode) int {
	return sumEvenGrandparentUtil(root, make(map[*TreeNode]*TreeNode), 0)
}
