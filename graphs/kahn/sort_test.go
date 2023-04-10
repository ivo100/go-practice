package main

import (
	"reflect"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	g := Graph{
		V:   4,
		Adj: make(map[int][]int),
	}
	g.addEdge(3, 2)
	g.addEdge(2, 1)
	g.addEdge(0, 1)
	expected := []int{0, 3, 2, 1}
	result := TopologicalSort(&g)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v but got %v", expected, result)
	}
}
