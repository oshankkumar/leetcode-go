package amountoftimeforbinarytreetobeinfected

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Vertex struct {
	Key           int
	AdjacencyList []*Vertex
}

type Graph struct {
	Vertices map[int]*Vertex
}

func (g *Graph) GetOrAddVertex(key int) *Vertex {
	v, ok := g.Vertices[key]
	if !ok {
		v = &Vertex{Key: key}
	}
	g.Vertices[key] = v
	return v
}

func (g *Graph) AddEdge(key1, key2 int) {
	v1, ok1 := g.Vertices[key1]
	v2, ok2 := g.Vertices[key2]
	if !ok1 || !ok2 {
		return
	}

	v1.AdjacencyList = append(v1.AdjacencyList, v2)
	v2.AdjacencyList = append(v2.AdjacencyList, v1)
}

func amountOfTime(root *TreeNode, start int) int {
	g := createGraph(root)
	infected := g.Vertices[start]

	visited := make(map[*Vertex]struct{})
	q := list.New()
	q.PushBack(infected)
	visited[infected] = struct{}{}

	var minutesElapsed int
	for q.Len() != 0 {
		qLen := q.Len()

		for i := 0; i < qLen; i++ {
			ver, _ := q.Remove(q.Front()).(*Vertex)

			for _, v := range ver.AdjacencyList {
				if _, ok := visited[v]; ok {
					continue
				}

				q.PushBack(v)
				visited[v] = struct{}{}
			}
		}

		minutesElapsed++
	}

	return minutesElapsed
}

func createGraph(root *TreeNode) *Graph {
	g := &Graph{
		Vertices: make(map[int]*Vertex),
	}

	q := list.New()
	q.PushBack(root)
	g.GetOrAddVertex(root.Val)

	for q.Len() != 0 {
		node, _ := q.Remove(q.Front()).(*TreeNode)

		if node.Left != nil {
			q.PushBack(node.Left)
			g.GetOrAddVertex(node.Left.Val)
			g.AddEdge(node.Val, node.Left.Val)
		}

		if node.Right != nil {
			q.PushBack(node.Right)
			g.GetOrAddVertex(node.Right.Val)
			g.AddEdge(node.Val, node.Right.Val)
		}
	}

	return g
}
