package topsort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	g := Graph{
		V:   6,
		Adj: make(map[int][]int),
	}
	g.AddEdge(5, 2)
	g.AddEdge(5, 0)
	g.AddEdge(4, 0)
	g.AddEdge(4, 1)
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)
	sorted := TopologicalSort(&g)
	//assert.EqualValues(t, []int{4, 5, 2, 0, 3, 1}, sorted)
	fmt.Println("Topologically sorted order:")
	for _, v := range sorted {
		fmt.Printf("%d ", v)
	}
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

func Test2(t *testing.T) {
	g := Graph{
		V:   9,
		Adj: make(map[int][]int),
	}

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
	sorted := TopologicalSort(&g)
	fmt.Println("Topologically sorted order:")
	for _, v := range sorted {
		fmt.Printf("%d ", v)
	}

}

func TestNotDAG(t *testing.T) {
	g := Graph{
		V:   3,
		Adj: make(map[int][]int),
	}
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	assert.Nil(t, TopologicalSort(&g))
}
