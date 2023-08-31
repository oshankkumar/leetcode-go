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

type Slice struct {
	arr []*TreeNode
}

func (s *Slice) Enqueue(node *TreeNode) error {
	s.arr = append(s.arr, node)
	return nil
}

func (s *Slice) Dequeue() (*TreeNode, error) {
	n := s.arr[0]
	s.arr = s.arr[1:]
	return n, nil
}

func (s *Slice) Len() int {
	return len(s.arr)
}

func rangeSumBSTBFS(root *TreeNode, low int, high int) int {
	var l Slice
	var sum int

	if err := l.Enqueue(root); err != nil {
		panic(err)
	}

	for l.Len() != 0 {
		node, err := l.Dequeue()
		if err != nil {
			panic(err)
		}

		if low <= node.Val && node.Val <= high {
			sum += node.Val
		}

		if node.Left != nil {
			if err := l.Enqueue(node.Left); err != nil {
				panic(err)
			}
		}

		if node.Right != nil {
			if err := l.Enqueue(node.Right); err != nil {
				panic(err)
			}
		}
	}

	return sum
}
