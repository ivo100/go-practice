package main

import "fmt"

type Graph struct {
	V   int
	Adj map[int][]int
}

func (g *Graph) addEdge(u, v int) {
	g.Adj[u] = append(g.Adj[u], v)
}

/*
In the TopologicalSort function, we first compute the in-degree of each vertex
in the graph, which is the number of incoming edges to the vertex.
We then initialize a queue q with all vertices that have zero in-degree.
We then repeatedly remove a vertex u from q, add it to the sorted order, and update the
in-degree of all its neighbors.
If a neighbor v of u now has zero in-degree, we add it to the queue q.
*/
func TopologicalSort(g *Graph) []int {
	inDegree := make(map[int]int)
	for u := 0; u < g.V; u++ {
		inDegree[u] = 0
	}
	for u := 0; u < g.V; u++ {
		for _, v := range g.Adj[u] {
			inDegree[v]++
		}
	}
	q := []int{}
	for u := 0; u < g.V; u++ {
		if inDegree[u] == 0 {
			q = append(q, u)
		}
	}
	sorted := []int{}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		sorted = append(sorted, u)
		for _, v := range g.Adj[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return sorted
}

func main() {
	g := Graph{
		V:   4,
		Adj: make(map[int][]int),
	}
	g.addEdge(3, 2)
	g.addEdge(2, 1)
	g.addEdge(0, 1)
	sorted := TopologicalSort(&g)
	fmt.Println("Topologically sorted order:")
	for _, v := range sorted {
		fmt.Printf("%d ", v)
	}
}
