package reverseoddlevelsofbinarytree

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	q := list.New()
	q.PushBack(root)
	currLevel := 0

	for q.Len() != 0 {
		currLen := q.Len()

		if currLevel&1 == 1 {
			left, right := q.Front(), q.Back()
			for {
				left.Value.(*TreeNode).Val, right.Value.(*TreeNode).Val = right.Value.(*TreeNode).Val, left.Value.(*TreeNode).Val
				if left.Next() == right {
					break
				}
				left = left.Next()
				right = right.Prev()
			}
		}

		for i := 0; i < currLen; i++ {
			node := q.Remove(q.Front()).(*TreeNode)

			if node.Left != nil {
				q.PushBack(node.Left)
			}

			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}

		currLevel++
	}

	return root
}

func reverseOddLevelsRec(root *TreeNode) *TreeNode {
	dfs(root.Left, root.Right, 1)
	return root
}

func dfs(left, right *TreeNode, level int) {
	if left == nil || right == nil {
		return
	}

	if level&1 == 1 {
		left.Val, right.Val = right.Val, left.Val
	}

	dfs(left.Left, right.Right, level+1)
	dfs(left.Right, right.Left, level+1)
}
