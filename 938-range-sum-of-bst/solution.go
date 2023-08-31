package rangesumofbst

import "errors"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	var val int
	if low <= root.Val && root.Val <= high {
		val = root.Val
	}

	return val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

func rangeSumBSTDFS(root *TreeNode, low int, high int) int {
	stk := []*TreeNode{root}
	var sum int

	for len(stk) != 0 {
		node := stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		if low <= node.Val && node.Val <= high {
			sum += node.Val
		}

		if node.Left != nil {
			stk = append(stk, node.Left)
		}

		if node.Right != nil {
			stk = append(stk, node.Right)
		}
	}

	return sum
}

var (
	ErrQueueEmpty = errors.New("queue is empty")
)

type Node struct {
	Data *TreeNode
	Next *Node
}

type List struct {
	Length int
	Head   *Node
	Tail   *Node
}

func (l *List) Enqueue(node *TreeNode) error {
	n := &Node{Data: node}

	if l.Head == nil {
		l.Head, l.Tail = n, n
		l.Length = 1
		return nil
	}

	l.Tail.Next = n
	l.Tail = n
	l.Length++
	return nil
}

func (l *List) Dequeue() (*TreeNode, error) {
	if l.Head == nil {
		return nil, ErrQueueEmpty
	}

	n := l.Head

	if l.Head == l.Tail {
		l.Head, l.Tail = nil, nil
		l.Length = 0
		return n.Data, nil
	}

	l.Head = l.Head.Next
	l.Length--

	return n.Data, nil
}

func (l *List) Len() int {
	return l.Length
}

type Queue []*TreeNode

func (q *Queue) Enqueue(node *TreeNode) error {
	*q = append(*q, node)
	return nil
}

func (q *Queue) Dequeue() (*TreeNode, error) {
	n := (*q)[0]
	*q = (*q)[1:]
	return n, nil
}

func (q *Queue) Len() int {
	return len(*q)
}

func rangeSumBSTBFS(root *TreeNode, low int, high int) int {
	var q Queue
	var sum int

	if err := q.Enqueue(root); err != nil {
		panic(err)
	}

	for q.Len() != 0 {
		node, err := q.Dequeue()
		if err != nil {
			panic(err)
		}

		if low <= node.Val && node.Val <= high {
			sum += node.Val
		}

		if node.Left != nil {
			if err := q.Enqueue(node.Left); err != nil {
				panic(err)
			}
		}

		if node.Right != nil {
			if err := q.Enqueue(node.Right); err != nil {
				panic(err)
			}
		}
	}

	return sum
}
