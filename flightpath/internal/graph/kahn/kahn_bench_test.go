package kahn

import (
	"testing"
)

// go test -bench=.

func BenchmarkTopologicalSort(b *testing.B) {
	//graph := NewGraph()
	// Add random edges to the graph
	for i := 0; i < 10000; i++ {
		//u := rand.Intn(string(g.V))
		//v := rand.Intn(g.V)
		//g.addEdge(u, v)
	}
	// Run the topological sort function b.N times
	//for n := 0; n < b.N; n++ {
	//	TopologicalSort(&g)
	//}
}
