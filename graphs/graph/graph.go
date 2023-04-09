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
	Color    string
	Distance uint
	Parent   *int
}

func NewNode(value string) *Node {
	return &Node{Value: value, Color: White, Distance: Infinity}
}

type Graph struct {
	V     int     // Number of vertices
	Nodes []Node  // Nodes
	Adj   [][]int // adj list - for each vertex - indices of connected nodes
	// store only node indices
	Q []int
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

func (g *Graph) init(start int) {
	// init
	g.Q = make([]int, 0)
	for i := 0; i < len(g.Nodes); i++ {
		n := &g.Nodes[i]
		if i == start {
			n.Distance = 0
			n.Parent = nil
			n.Color = Gray
			// enqueue start node
			g.Q = append(g.Q, i)
			continue
		}
		n.Color = White
		n.Distance = Infinity
		n.Parent = nil
	}
}

func (g *Graph) BFS(start int) {
	g.init(start)
	for len(g.Q) > 0 {
		i := g.Q[0]
		u := &g.Nodes[i]
		g.Q = g.Q[1:]
		for j := 0; j < len(g.Adj[i]); j++ {
			v := &g.Nodes[g.Adj[i][j]]
			if v.Color == White {
				v.Color = Gray
				v.Distance = u.Distance + 1
				v.Parent = &u.Idx
				// enqueue graph node
				g.Q = append(g.Q, v.Idx)
			}
		}
		u.Color = Black
		//log.Printf("Graph:\n%v", g.String())
	}
}

func (g *Graph) String() string {
	var b strings.Builder
	for i, node := range g.Nodes {
		b.WriteString(fmt.Sprintf("%s (%v) -> ", node.Value, node.Color))
		for j := 0; j < len(g.Adj[i]); j++ {
			if j > 0 {
				b.WriteString(", ")
			}
			a := g.Adj[i][j]
			n := g.Nodes[a]
			b.WriteString(fmt.Sprintf("%s (%v %v)", n.Value, n.Color, n.Distance))
			if n.Parent != nil {
				b.WriteString(fmt.Sprintf(" parent %v", *n.Parent))
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}
