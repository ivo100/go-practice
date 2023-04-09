package bfs

import (
	"log"
	"testing"
)

func TestBFS(t *testing.T) {
	g := NewGraph(6)
	g.AddNode(0, "A")
	g.AddNode(1, "B")
	g.AddNode(2, "C")
	g.AddNode(3, "D")
	g.AddNode(4, "E")
	g.AddNode(5, "F")
	// A -> B, D
	g.AddEdge(0, 1)
	g.AddEdge(0, 3)
	// B -> E
	g.AddEdge(1, 4)
	// D -> B
	g.AddEdge(3, 1)
	// E -> C
	g.AddEdge(4, 2)
	// F -> C
	g.AddEdge(5, 2)

	log.Printf("Graph before BFS:\n%v", g.String())

	g.BFS(0)

	log.Printf("Graph after BFS:\n%v", g.String())
}
