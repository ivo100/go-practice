package kahn

import (
	"github.com/dominikbraun/graph"
	"log"
)

type Node struct {
	ID   int
	Name string
}

type Graph struct {
	nodes []*Node
	edges map[int][]int
}

// NewGraph creates a new graph
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

// TopologicalSort performs the sorting algorithm
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
	// this linear search is not good for large graphs
	// can be replaced with a more efficient hashmap
	for _, node := range g.nodes {
		if node.ID == id {
			return node
		}
	}
	return nil
}

func (g *Graph) findNodeByName(name string) *Node {
	// this linear search is not good for large graphs
	// can be replaced with a more efficient hashmap
	for _, node := range g.nodes {
		if node.Name == name {
			return node
		}
	}
	return nil
}

// Sort performs topological sorting of DAG defined by vertices and edges
func Sort(vertices []string, edges [][]string) ([]string, error) {
	g := NewGraph()
	for idx, v := range vertices {
		node := &Node{ID: idx, Name: v}
		g.AddNode(node)
	}
	for _, e := range edges {
		n1 := g.findNodeByName(e[0])
		if n1 == nil {
			return nil, graph.ErrEdgeNotFound
		}
		n2 := g.findNodeByName(e[1])
		if n1 == nil {
			return nil, graph.ErrEdgeNotFound
		}
		g.AddEdge(n1.ID, n2.ID)
	}
	nodes := g.TopologicalSort()
	res := make([]string, 0)
	for _, n := range nodes {
		res = append(res, n.Name)
	}
	return res, nil
}
