package graph

import (
	"log"
	"testing"
)

func TestGraph_Add(t *testing.T) {
	g := NewGraph(6)
	g.AddNode(0, "A")
	g.AddNode(1, "B")
	g.AddNode(2, "C")
	g.AddNode(3, "D")
	g.AddNode(4, "E")
	g.AddNode(5, "F")
	g.AddEdge(0, 1)
	g.AddEdge(0, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	g.AddEdge(3, 1)

	log.Printf("Graph:\n%v", g.String())
}
