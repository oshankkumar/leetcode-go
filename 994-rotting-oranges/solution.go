package rottingoranges

import "container/list"

const (
	Empty int = iota
	Fresh
	Rotten
)

type Vertex struct {
	Key           int
	Val           int
	AdjacencyList []*Vertex
}

type Graph struct {
	Vertices map[int]*Vertex
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

func (g *Graph) BFSMinDuration(totalFresh int, rottenQ *list.List) int {
	var minutesElapsed int
	visited := make(map[*Vertex]struct{})

	for totalFresh > 0 && rottenQ.Len() != 0 {
		qLen := rottenQ.Len()
		minutesElapsed++
		for i := 0; i < qLen; i++ {
			ver := rottenQ.Remove(rottenQ.Front()).(*Vertex)

			for _, v := range ver.AdjacencyList {
				if _, ok := visited[v]; ok {
					continue
				}
				visited[v] = struct{}{}
				if v.Val == Fresh {
					v.Val = Rotten
					totalFresh--
					rottenQ.PushBack(v)
				}
			}
		}
	}

	if totalFresh > 0 {
		return -1
	}

	return minutesElapsed
}

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	pos := func(i, j int) int {
		return (i * n) + j
	}

	validIndex := func(i, j int) bool {
		return i <= m-1 && j <= n-1
	}

	gr := &Graph{Vertices: make(map[int]*Vertex)}
	var freshCount int
	rottenQ := list.New()

	for i, g := range grid {
		for j, cell := range g {
			if cell == Empty {
				continue
			}

			if cell == Fresh {
				freshCount++
			}

			v := &Vertex{Key: pos(i, j), Val: cell}

			if v.Val == Rotten {
				rottenQ.PushBack(v)
			}

			gr.Vertices[v.Key] = v
		}
	}

	for i, row := range grid {
		for j, c := range row {
			if c == Empty {
				continue
			}

			if validIndex(i, j+1) && grid[i][j+1] != Empty {
				gr.AddEdge(pos(i, j), pos(i, j+1))
			}

			if validIndex(i+1, j) && grid[i+1][j] != Empty {
				gr.AddEdge(pos(i, j), pos(i+1, j))
			}
		}
	}

	return gr.BFSMinDuration(freshCount, rottenQ)
}
