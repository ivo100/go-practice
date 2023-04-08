package graph

import (
	"fmt"
	"strings"
)

const (
	White    = "white"
	Black    = "black"
	Gray     = "gray"
	Infinity = 999999999
)

type Node struct {
	Idx      int
	Value    string
	Parent   *Node
	Distance uint
	Color    string
}

func NewNode(value string) *Node {
	return &Node{Value: value, Color: White, Distance: Infinity}
}

type Graph struct {
	V     int     // Number of vertices
	Nodes []Node  // Nodes
	Adj   [][]int // adj list - for each vertex - indices of connected nodes
}

func NewGraph(numNodes int) *Graph {
	g := &Graph{
		V:     numNodes,
		Nodes: make([]Node, numNodes),
		Adj:   make([][]int, numNodes),
	}
	for i := 0; i < numNodes; i++ {
		g.Adj[i] = make([]int, 0)
	}
	return g
}

func (g *Graph) AddNode(idx int, value string) {
	node := NewNode(value)
	node.Idx = idx
	g.Nodes[idx] = *node
}

func (g *Graph) AddEdge(from int, to int) {
	g.Adj[from] = append(g.Adj[from], to)
}

func (g *Graph) String() string {
	var b strings.Builder
	for i, node := range g.Nodes {
		b.WriteString(fmt.Sprintf("%s -> ", node.Value))
		for j := 0; j < len(g.Adj[i]); j++ {
			if j > 0 {
				b.WriteString(", ")
			}
			a := g.Adj[i][j]
			b.WriteString(fmt.Sprintf("%s", g.Nodes[a].Value))
		}
		b.WriteString("\n")
	}
	return b.String()
}
