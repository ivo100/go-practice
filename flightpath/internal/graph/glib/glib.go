package glib

import (
	"github.com/dominikbraun/graph"
)

// Sort performs topological sorting of DAG using graph library and string hash / string values = airport codes
func Sort(vertices []string, edges [][]string) ([]string, error) {
	g := graph.New(graph.StringHash, graph.Directed())

	for _, v := range vertices {
		if err := g.AddVertex(v); err != nil {
			return nil, err
		}
	}
	for _, e := range edges {
		err := g.AddEdge(e[0], e[1])
		if err != nil {
			return nil, err
		}
	}
	result, err := graph.TopologicalSort(g)
	if err != nil {
		return nil, err
	}
	return result, nil
}
