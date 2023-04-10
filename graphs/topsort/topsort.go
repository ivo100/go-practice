package main

import (
	"fmt"
	"log"
)

type Node struct {
	ID    int
	Label string
	Name  string
}

type Graph struct {
	nodes []*Node
	edges map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		nodes: []*Node{},
		edges: make(map[int][]int),
	}
}

func (g *Graph) AddNode(node *Node) {
	g.nodes = append(g.nodes, node)
}

func (g *Graph) AddEdge(nodeID1, nodeID2 int) {
	g.edges[nodeID1] = append(g.edges[nodeID1], nodeID2)
}

func (g *Graph) TopologicalSort() []*Node {
	// 1. Calculate in-degree of all nodes
	inDegree := make(map[int]int)
	for _, node := range g.nodes {
		inDegree[node.ID] = 0
	}
	for _, neighbors := range g.edges {
		for _, neighborID := range neighbors {
			inDegree[neighborID]++
		}
	}

	// 2. Create a queue and enqueue all nodes with in-degree 0
	queue := []*Node{}
	for _, node := range g.nodes {
		if inDegree[node.ID] == 0 {
			queue = append(queue, node)
		}
	}

	// 3. Process the queue
	sortedNodes := []*Node{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sortedNodes = append(sortedNodes, node)

		for _, neighborID := range g.edges[node.ID] {
			inDegree[neighborID]--
			if inDegree[neighborID] == 0 {
				queue = append(queue, g.findNodeByID(neighborID))
			}
		}
	}

	// 4. Check for cycles
	if len(sortedNodes) != len(g.nodes) {
		log.Printf("Cycle detected")
		return []*Node{} // There is a cycle
	}

	return sortedNodes
}

func (g *Graph) findNodeByID(id int) *Node {
	for _, node := range g.nodes {
		if node.ID == id {
			return node
		}
	}
	return nil
}

const (
	Pants     = 1
	Belt      = 2
	Shirt     = 3
	Socks     = 4
	Shoes     = 5
	Underwear = 6
)

func main() {
	graph := NewGraph()

	pants := &Node{ID: Pants, Label: "A", Name: "pants"}
	belt := &Node{ID: Belt, Label: "B", Name: "belt"}
	shirt := &Node{ID: Shirt, Label: "C", Name: "shirt"}
	socks := &Node{ID: Socks, Label: "D", Name: "socks"}
	shoes := &Node{ID: Shoes, Label: "E", Name: "shoes"}
	underwear := &Node{ID: Underwear, Label: "F", Name: "underwear"}

	graph.AddNode(pants)
	graph.AddNode(belt)
	graph.AddNode(shirt)
	graph.AddNode(socks)
	graph.AddNode(shoes)
	graph.AddNode(underwear)

	graph.AddEdge(Underwear, Shirt)
	graph.AddEdge(Underwear, Socks)
	graph.AddEdge(Underwear, Pants)
	graph.AddEdge(Pants, Belt)
	graph.AddEdge(Socks, Shoes)
	graph.AddEdge(Shirt, Belt)

	//for _, node := range graph.nodes {
	//	fmt.Printf("Node ID: %d, Label: %s, Name: %s\n", node.ID, node.Label, node.Name)
	//}

	sortedNodes := graph.TopologicalSort()
	for _, node := range sortedNodes {
		fmt.Printf("Node ID: %d, Label: %s, Name: %s\n", node.ID, node.Label, node.Name)
	}
	// Output:
	// Node ID: 1, Label: A, Name: Node A
	// Node ID: 2, Label: B, Name: Node B
	// Node ID: 3, Label:
}

/***
package topsort

// ChatGPT version - seems the best

type Graph struct {
	V   int
	Adj map[int][]int
}

func (g *Graph) AddEdge(u, v int) {
	g.Adj[u] = append(g.Adj[u], v)
}

// TopologicalSort sorts given DAG so no vertex appears before any of its descendants
// if there are cycles (input is not a DAG) - returns nil slice
func TopologicalSort(dag *Graph) []int {
	//compute the in-degree of each vertex in the graph, which is the number of
	//incoming edges to the vertex.
	inDegree := make(map[int]int)
	for u := 0; u < dag.V; u++ {
		inDegree[u] = 0
	}
	for u := 0; u < dag.V; u++ {
		for _, v := range dag.Adj[u] {
			inDegree[v]++
		}
	}
	// initialize queue q with all vertices that have zero incoming.
	// repeatedly remove vertex u from q, add it to the result in order,
	// decrement  the in-degree of all its neighbors.
	// If a neighbor v of u has zero in-degree, add it to the queue q.
	var q []int
	for u := 0; u < dag.V; u++ {
		if inDegree[u] == 0 {
			q = append(q, u)
		}
	}
	var result []int
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, v := range dag.Adj[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return result
}
***/
