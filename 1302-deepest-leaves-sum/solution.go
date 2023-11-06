package deepestleavessum

func deepestLeavesSum(root *TreeNode) int {
	var (
		q   []*TreeNode
		sum int
	)

	if root != nil {
		q = append(q, root)
	}

	for {
		qLen := len(q)
		var levelSum int
		
		for i := 0; i < qLen; i++ {
			front := q[0]
			q = q[1:]

			levelSum += front.Val

			if front.Left != nil {
				q = append(q, front.Left)
			}

			if front.Right != nil {
				q = append(q, front.Right)
			}
		}

		if len(q) == 0 {
			sum = levelSum
			break
		}
	}

	return sum
}
