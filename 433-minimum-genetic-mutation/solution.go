package minimumgeneticmutation

import "container/list"

type Vertex struct {
	Key           string
	AdjacencyList []*Vertex
}

type Graph struct {
	Vertices map[string]*Vertex
}

func NewGraph(keys ...string) *Graph {
	g := &Graph{Vertices: make(map[string]*Vertex)}
	for _, k := range keys {
		g.Vertices[k] = &Vertex{Key: k}
	}
	return g
}

func (g *Graph) AddEdge(key1, key2 string) {
	v1, ok1 := g.Vertices[key1]
	v2, ok2 := g.Vertices[key2]
	if !ok1 || !ok2 {
		return
	}
	v1.AdjacencyList = append(v1.AdjacencyList, v2)
	v2.AdjacencyList = append(v2.AdjacencyList, v1)
}

func (g *Graph) BFS(srcKey string, destKey string) int {
	visited := make(map[*Vertex]struct{})
	q := list.New()

	srcVer := g.Vertices[srcKey]
	q.PushBack(srcVer)
	visited[srcVer] = struct{}{}

	var dist int

	for q.Len() != 0 {
		qLen := q.Len()
		for i := 0; i < qLen; i++ {
			front := q.Remove(q.Front()).(*Vertex)
			if front.Key == destKey {
				return dist
			}
			for _, ver := range front.AdjacencyList {
				if _, ok := visited[ver]; ok {
					continue
				}

				visited[ver] = struct{}{}
				q.PushBack(ver)
			}
		}

		dist++
	}

	return -1
}

func minMutation(startGene string, endGene string, bank []string) int {
	validGenes := make([]string, len(bank))
	copy(validGenes, bank)
	validGenes = append(validGenes, startGene)

	g := NewGraph(validGenes...)
	for i := 0; i < len(validGenes); i++ {
		for j := i + 1; j < len(validGenes); j++ {
			if validMutation(validGenes[i], validGenes[j]) {
				g.AddEdge(validGenes[i], validGenes[j])
			}
		}
	}

	return g.BFS(startGene, endGene)
}

func validMutation(g1, g2 string) bool {
	const maxDiff = 1

	var diff int
	for i := 0; i < len(g1); i++ {
		if g1[i] != g2[i] {
			diff++
		}
		if diff > maxDiff {
			return false
		}
	}

	return diff <= maxDiff
}
