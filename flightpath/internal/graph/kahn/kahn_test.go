package kahn

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

const (
	Pants     = 1
	Belt      = 2
	Shirt     = 3
	Socks     = 4
	Shoes     = 5
	Underwear = 6
)

func Test_KahnTopologicalSort(t *testing.T) {
	graph := NewGraph()

	pants := &Node{ID: Pants, Name: "pants"}
	belt := &Node{ID: Belt, Name: "belt"}
	shirt := &Node{ID: Shirt, Name: "shirt"}
	socks := &Node{ID: Socks, Name: "socks"}
	shoes := &Node{ID: Shoes, Name: "shoes"}
	underwear := &Node{ID: Underwear, Name: "underwear"}

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

	s := graph.TopologicalSort()
	assert.Equal(t, 6, len(s))
	assert.Equal(t, "underwear", s[0].Name)
	assert.Equal(t, "belt", s[5].Name)

}

func Test_Sort(t *testing.T) {
	nodes := []string{"IND", "EWR", "SFO", "ATL", "GSO"}
	edges := [][]string{
		{"IND", "EWR"},
		{"SFO", "ATL"},
		{"GSO", "IND"},
		{"ATL", "GSO"},
	}
	p, err := Sort(nodes, edges)
	assert.NoError(t, err)
	log.Printf("%v", p)
	require.Equal(t, 5, len(p))
	assert.Equal(t, "SFO", p[0])
	assert.Equal(t, "EWR", p[4])
}
