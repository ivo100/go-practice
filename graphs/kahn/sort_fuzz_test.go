//go:build gofuzz
// +build gofuzz

// go test -fuzz=Fuzz -fuzztime=10s
package main

import "testing/fuzz"

func Fuzz(data []byte) int {
	// Parse input data into graph
	var g Graph
	err := fuzz.Unmarshal(data, &g)
	if err != nil {
		return -1
	}

	// Run topological sort
	result := TopologicalSort(&g)

	// Return 0 if sort succeeds, -1 otherwise
	if isTopologicallySorted(g, result) {
		return 0
	} else {
		return -1
	}
}
