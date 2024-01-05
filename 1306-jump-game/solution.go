package jumpgame

import "container/list"

type Vertex struct {
	key           int
	value         int
	adjacencyList []*Vertex
}

func NewVertex(k, v int) *Vertex {
	return &Vertex{key: k, value: v}
}

type Graph struct {
	vertices map[int]*Vertex
}

func NewGraph(vv ...*Vertex) *Graph {
	g := &Graph{vertices: make(map[int]*Vertex)}
	for _, v := range vv {
		g.vertices[v.key] = v
	}
	return g
}

func (g *Graph) AddEdge(k1, k2 int) {
	v1, ok1 := g.vertices[k1]
	v2, ok2 := g.vertices[k2]
	if !ok1 || !ok2 {
		return
	}
	v1.adjacencyList = append(v1.adjacencyList, v2)
}

func (g *Graph) BFS(srcKey int, destVal int) bool {
	visited := make(map[*Vertex]struct{})

	q := list.New()
	q.PushBack(g.vertices[srcKey])

	for q.Len() != 0 {
		v, _ := q.Remove(q.Front()).(*Vertex)
		if v.value == destVal {
			return true
		}

		for _, ver := range v.adjacencyList {
			_, ok := visited[ver]
			if ok {
				continue
			}
			visited[ver] = struct{}{}
			q.PushBack(ver)
		}
	}

	return false
}

func canReach(arr []int, start int) bool {
	vv := make([]*Vertex, len(arr))
	for i, val := range arr {
		vv[i] = NewVertex(i, val)
	}

	g := NewGraph(vv...)

	for i, val := range arr {
		g.AddEdge(i, i+val)
		g.AddEdge(i, i-val)
	}

	return g.BFS(start, 0)
}
