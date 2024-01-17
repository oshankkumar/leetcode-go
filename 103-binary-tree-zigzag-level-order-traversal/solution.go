package binary_tree_zigzag_level_order_traversal

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	dendQ := list.New()
	currDir := 0x1

	if root != nil {
		dendQ.PushBack(root)
	}

	var result [][]int
	for dendQ.Len() != 0 {
		qLen := dendQ.Len()

		var levelNodes []int
		for i := 0; i < qLen; i++ {
			if currDir == 0x1 {
				node, _ := dendQ.Remove(dendQ.Front()).(*TreeNode)
				levelNodes = append(levelNodes, node.Val)

				if node.Left != nil {
					dendQ.PushBack(node.Left)
				}

				if node.Right != nil {
					dendQ.PushBack(node.Right)
				}
			} else {
				node, _ := dendQ.Remove(dendQ.Back()).(*TreeNode)
				levelNodes = append(levelNodes, node.Val)

				if node.Right != nil {
					dendQ.PushFront(node.Right)
				}

				if node.Left != nil {
					dendQ.PushFront(node.Left)
				}

			}
		}

		result = append(result, levelNodes)
		currDir ^= 0x1
	}

	return result
}
