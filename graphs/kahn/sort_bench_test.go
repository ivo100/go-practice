package main

import (
	"math/rand"
	"testing"
)

// go test -bench=.

func BenchmarkTopologicalSort(b *testing.B) {
	g := Graph{
		V:   1000,
		Adj: make(map[int][]int),
	}
	// Add random edges to the graph
	for i := 0; i < 10000; i++ {
		u := rand.Intn(g.V)
		v := rand.Intn(g.V)
		g.addEdge(u, v)
	}
	// Run the topological sort function b.N times
	for n := 0; n < b.N; n++ {
		TopologicalSort(&g)
	}
}
