package dfs

import (
	"log"
	"testing"
)

func TestDFS(t *testing.T) {
	g := NewGraph(6)
	g.AddNode(0, "U")
	g.AddNode(1, "V")
	g.AddNode(2, "W")
	g.AddNode(3, "X")
	g.AddNode(4, "Y")
	g.AddNode(5, "Z")
	// u -> v, x
	g.AddEdge(0, 1)
	g.AddEdge(0, 3)
	// v -> y
	g.AddEdge(1, 4)
	// w -> y, z
	g.AddEdge(2, 4)
	g.AddEdge(2, 5)
	// x -> v
	g.AddEdge(3, 1)
	// y -> x
	g.AddEdge(4, 3)
	// z -> z
	g.AddEdge(5, 5)

	log.Printf("Graph before DFS:\n%v", g.String())

	g.DFS()

	log.Printf("Graph after DFS:\n%v", g.String())
}

const (
	Under  = 0
	Pants  = 1
	Shirt  = 2
	Tie    = 3
	Jacket = 4
	Belt   = 5
	Socks  = 6
	Shoes  = 7
	Watch  = 8
)

func TestTopoSort(t *testing.T) {
	g := NewGraph(9)

	// 9 nodes
	g.AddNode(Under, "under")
	g.AddNode(Pants, "pants")
	g.AddNode(Shirt, "shirt")
	g.AddNode(Tie, "tie")
	g.AddNode(Jacket, "jacket")
	g.AddNode(Belt, "belt")
	g.AddNode(Socks, "socks")
	g.AddNode(Shoes, "shoes")
	g.AddNode(Watch, "watch")

	// 9 edges
	g.AddEdge(Under, Pants)
	g.AddEdge(Under, Shoes)
	g.AddEdge(Pants, Belt)
	g.AddEdge(Belt, Jacket)
	g.AddEdge(Shirt, Tie)
	g.AddEdge(Shirt, Belt)
	g.AddEdge(Tie, Jacket)
	g.AddEdge(Socks, Shoes)
	g.AddEdge(Pants, Shoes)

	log.Printf("Graph before DFS:\n%v", g.String())

	g.DFS()

	log.Printf("Graph after DFS:\n%v", g.String())

	log.Printf("Sorted %v", g.Sorted)
	for i := 0; i < len(g.Sorted); i++ {
		log.Printf("%v", g.Nodes[i].Value)
	}

}
