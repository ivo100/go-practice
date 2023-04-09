package dfs

import (
	"fmt"
	"strings"
)

const (
	White = "white"
	Black = "black"
	Gray  = "gray"
)

type Node struct {
	Idx    int
	Value  string
	Color  string
	Enter  int
	Exit   int
	Parent *int
}

func NewNode(value string) *Node {
	return &Node{Value: value, Color: White}
}

type Graph struct {
	V      int     // Number of vertices
	Nodes  []Node  // Nodes
	Adj    [][]int // adj list - for each vertex - indices of connected nodes
	time   int
	Sorted []int
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

func (g *Graph) init() {
	g.time = 0
	g.Sorted = make([]int, 0)
	for i := 0; i < len(g.Nodes); i++ {
		n := &g.Nodes[i]
		n.Color = White
		n.Parent = nil
	}
}

func (g *Graph) DFS() {
	g.init()
	for i := 0; i < len(g.Nodes); i++ {
		n := &g.Nodes[i]
		if n.Color == White {
			g.Visit(n)
		}
	}
}

func (g *Graph) Visit(u *Node) {
	g.time++
	u.Enter = g.time
	u.Color = Gray
	i := u.Idx
	// explore edge (u, v)
	for j := 0; j < len(g.Adj[i]); j++ {
		v := &g.Nodes[g.Adj[i][j]]
		if v.Color == White {
			v.Parent = &u.Idx
			g.Visit(v)
		}
	}
	u.Color = Black
	g.time++
	u.Exit = g.time
	g.Sorted = append(g.Sorted, u.Idx)
	//log.Printf("Graph:\n%v", g.String())
}

func (g *Graph) String() string {
	var b strings.Builder
	for i, node := range g.Nodes {
		b.WriteString(fmt.Sprintf("%s (%v %v/%v) -> ", node.Value, node.Color, node.Enter, node.Exit))
		for j := 0; j < len(g.Adj[i]); j++ {
			if j > 0 {
				b.WriteString(", ")
			}
			a := g.Adj[i][j]
			n := g.Nodes[a]
			b.WriteString(fmt.Sprintf("%s (%v %v/%v)", n.Value, n.Color, n.Enter, n.Exit))
			if n.Parent != nil {
				b.WriteString(fmt.Sprintf(" parent %v", *n.Parent))
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}
