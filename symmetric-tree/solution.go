package symmetrictree

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricTreesRecursion(root.Left, root.Right)
}

func isSymmetricTreesRecursion(r1, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}

	if r1 == nil || r2 == nil {
		return false
	}

	return (r1.Val == r2.Val) &&
		isSymmetricTreesRecursion(r1.Left, r2.Right) &&
		isSymmetricTreesRecursion(r1.Right, r2.Left)
}

func isSymmetricTreesBFS(root *TreeNode) bool {
	var queue *list.List

	queue.PushBack(root.Left)
	queue.PushBack(root.Right)

	for queue.Len() != 0 {
		n1, _ := queue.Remove(queue.Front()).(*TreeNode)
		n2, _ := queue.Remove(queue.Front()).(*TreeNode)

		if n1 == nil && n2 == nil {
			continue
		}

		if n1 == nil || n2 == nil {
			return false
		}

		if n1.Val != n2.Val {
			return false
		}

		queue.PushBack(n1.Left)
		queue.PushBack(n2.Right)
		queue.PushBack(n1.Right)
		queue.PushBack(n2.Left)
	}

	return true
}
