package binarytreelevelordertraversal

type Q []*TreeNode

func (q *Q) PushBack(n *TreeNode) {
	*q = append(*q, n)
}

func (q *Q) PopFront() (*TreeNode, bool) {
	if len(*q) == 0 {
		return nil, false
	}

	n := (*q)[0]
	*q = (*q)[1:]

	return n, true
}

func (q Q) Len() int {
	return len(q)
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int

	var q Q
	if root != nil {
		q.PushBack(root)
	}

	for q.Len() != 0 {
		currLen := q.Len()
		var nodes []int
		for i := 0; i < currLen; i++ {
			n, _ := q.PopFront()
			nodes = append(nodes, n.Val)

			if n.Left != nil {
				q.PushBack(n.Left)
			}
			if n.Right != nil {
				q.PushBack(n.Right)
			}
		}
		res = append(res, nodes)
	}

	return res
}
